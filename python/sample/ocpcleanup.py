#!/usr/bin/env python

import sys
import os
import re

if __name__ == "__main__":

    f = open("config.json", "r")
    contents_str = f.read()

    contents_new_str = re.sub(r"ocp-.*\"", "ocp-" + sys.argv[1] + "\"", contents_str)
    f1 = open("config1.json", "w")
    f1.write(contents_new_str)
    f1.close()

    os.system("bazelisk run cleanupcluster deleteCluster -- --config=" + os.path.abspath(f1.name))
