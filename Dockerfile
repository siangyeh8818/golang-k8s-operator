FROM siangyeh8818/go-base-environment:1.12

#RUN go get -u -v github.com/kardianos/govendor

COPY golang-k8s-operator go/src/golang-k8s-operator
RUN govendor fetch github.com/kubernetes/client-go
