from __future__ import division
import torch
import torch.nn as nn
from PIL import Image
import boto3
import pickle
import settings
from os import listdir
from os.path import isfile, join
import re

from torchvision import datasets, models, transforms


class S3Access:
    def __init__(self):
        self.aws_access_key_id = settings.S3_MODELS_ACCESS_KEY
        self.aws_secret_access_key = settings.S3_MODELS_SECRET_KEY
        self.s3_bucket_name = settings.S3_MODELS_BUCKET
        self.s3_region_name = settings.S3_MODELS_REGION

        # self.aws_access_key_id = "AKIAURB4OY5SOGIFG7UG"
        # self.aws_secret_access_key = "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"
        # self.s3_bucket_name = "manager-classifier2"
        # self.s3_region_name = "eu-west-1"

    def download_model(self, path_current: str = 'active_models/listing_image_classifier') -> dict:
        s3_resource = boto3.resource('s3', aws_access_key_id=self.aws_access_key_id,
                                     aws_secret_access_key=self.aws_secret_access_key)
        bucket_ = s3_resource.Bucket(self.s3_bucket_name)
        collection = bucket_.objects.filter(Prefix=path_current).all()
        state_dict = None
        if collection:
            objects = list(collection)
            state_dict = pickle.loads(bucket_.Object(objects[0].key).get()['Body'].read())
        return state_dict


class ToiletFinder:
    def __init__(self, predict=False):
        self.labels = ['bathroom',
                       'bedroom',
                       'exterior',
                       'interior',
                       'kitchen',
                       'living_room',
                       'pool']
        self.toilet_index = 0
        self.model_location = ".cache/listing_image_classifier"
        self.model_name = "resnet"
        # Number of classes in the dataset
        self.num_classes = 7
        # Flag for feature extracting. When False, we finetune the whole model,
        #   when True we only update the reshaped layer params
        self.feature_extract = True
        self.input_size = 0  # image dim
        self.device = None
        self.model = self.load_model()

    @staticmethod
    def set_parameter_requires_grad(model, feature_extracting):
        if feature_extracting:
            for param in model.parameters():
                param.requires_grad = False

    @staticmethod
    def pre_process(image: Image.Image, input_size: int):
        return transforms.Compose([
                transforms.Resize(input_size),
                transforms.CenterCrop(input_size),
                transforms.ToTensor(),
                transforms.Normalize([0.485, 0.456, 0.406], [0.229, 0.224, 0.225])
            ])(image)

    def initialize_model(self, model_name, num_classes, feature_extract):
        # Initialize these variables which will be set in this if statement. Each of these
        #   variables is model specific.
        self.device = torch.device("cuda:0" if torch.cuda.is_available() else "cpu")

        model_ft = models.resnet18(pretrained=True)
        self.set_parameter_requires_grad(model_ft, feature_extract)
        num_ftrs = model_ft.fc.in_features
        # activation function ?
        model_ft.fc = nn.Linear(num_ftrs, num_classes)
        input_size = 224

        self.input_size = input_size
        print("Model initialized")


        return model_ft, input_size

    # try to do it in the beginning in the file, double if check if local i used somewhere
    def load_model(self):
        self.model, self.input_size = self.initialize_model(self.model_name, self.num_classes, self.feature_extract)
        if isfile(self.model_location):
            # model = open(self.model_location, "rb")
            # print("loading model locally")
            # state_dict = pickle.load(model)
            # model.close()
            with open(self.model_location, "rb") as model:
                print("loading model locally")
                state_dict = pickle.load(model)
        else:
            s3 = S3Access()
            state_dict = s3.download_model()
            print("Retrieving model from S3")
            with open(self.model_location, 'wb') as model_file:
                pickle.dump(state_dict, model_file)
        self.model.load_state_dict(state_dict)
        return self.model

    def save_model(self):
        s3 = S3Access()
        s3.model_to_s3(self.model)

    def predict(self, image_file):
        # preprocessing
        # transform = self.pre_process['val']
        image = Image.open(image_file)
        if (len(image.getbands()) > 3):
            # RGBA does not work. Must convert back to RGB. Some data loss will take place in this process
            image = image.convert("RGB")
        # apply pre-processing
        img_t = self.pre_process(image, self.input_size)
        # adapting above output to model input (shape of matrix)
        batch_t = torch.unsqueeze(img_t, 0)
        # loading model every prediction (Ricardo doesnt like)
        self.model.eval()
        out = self.model(batch_t)
        _, indices = torch.sort(out, descending=True)
        percentage = torch.nn.functional.softmax(out, dim=1)[0] * 100
        return [(self.labels[idx], percentage[idx].item()) for idx in indices[0][:3]]
