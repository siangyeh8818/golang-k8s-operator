# golang-k8s-operator

#用途
將k8s現有環境轉化成部署的描述檔(可以想成gdeyamlOperator的部署規格)

gdeyamOperator的repo : https://github.com/siangyeh8818/gdeyamlOperator

#限制
受到執行的host上面的kubeconfig權限所規範

# 用法
./k8sclone --o owner_deploy.yml --namespace workflow-stable

| flag      | 說明    | 預設值     |
| ---------- | :-----------:  | :-----------: |
|   o   | 輸出的yml檔名   | deploy.yml   |
|  namespace    |  要將哪個namespace的資源輸出   | default    |
