FROM centos:7
ADD ./main  /main
ADD ./index.html /index.html
CMD chmod +x  /main

