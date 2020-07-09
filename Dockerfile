FROM ubuntu:18.04
ADD ./main  /main
ADD ./index.html /index.html
CMD /main

