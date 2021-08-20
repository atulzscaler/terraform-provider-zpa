""" Script to update vendor dependencies
    Needs: sudo -H pip install -U checksumdir  """
import os
import shutil
import json
import subprocess
import checksumdir
TERRAFORM_VERSION = "v0.11.11"


def find(name, path):
    """ Find a directory in path"""
    result = []
    for root, dirs, files in os.walk(path):
        if name in dirs:
            result.append(os.path.join(root, name))
    return result


def get_immediate_subdirectories(a_dir):
    """ Get immediate subdirectories"""
    return [name for name in os.listdir(a_dir) if os.path.isdir(os.path.join(a_dir, name))]