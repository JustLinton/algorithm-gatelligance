from fileinput import filename
import sys
import os

#任务id，bili视频链接
fileName = "./tmp/"+sys.argv[1]+".mp3"
# print(fileName+" "+sys.argv[2])
# print("donwload...")
retVal1 = os.system("python3 ./pyScripts/dnld.py "+fileName+" "+sys.argv[2])
if retVal1!=0:
    print('download error.')
    exit(0)
# print(retVal1)
# print("process...")
retVal2 = os.system("python3 ./pyScripts/xfr.py "+fileName)
if retVal2!=0:
    print('xfr error.')
    exit(0)
# print(retVal2)