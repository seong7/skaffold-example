**[Skaffold](https://skaffold.dev/docs/) 는 Kubenetes 환경에서 Continuous Build & Deploy 를 간편하게 도와주는 Tool 입니다.**

## Example 실행법 :

> k8s cluster 가 구축되어 있고 kubectl 이 해당 cluster 를 잘 바라보고 있음을 가정합니다.

Skaffold 를 설치합니다.

[examples/getting-started](./examples/getting-started) 로 이동

1. `skaffold.yaml` 에서 `build.artifacts[].image` 에 docker image 를 push 할 Container Registry 와 image name 입력
    - default 는 docker registry
    - Github Container Registry 나 private registry 를 이용중이라면 접속 정보를 담은 [secret](https://kubernetes.io/ko/docs/concepts/configuration/secret/) 이 k8s cluster 내 생성되어 있어야 합니다.

2. `k8s-pod.yaml` 에서 `spec.containers[].image` 에 위에서 입력한 문자열과 동일하게 입력 (tag 는 입력하지 않아도 됨)

3. skaffold 실행 (watch 모드) 
 ```bash
skaffold dev
```

4. local 에서 `main.go` 를 수정하여 Skaffold 가 자동으로 아래 과정을 실행하는 것을 확인합니다. 
  - Image build (using cache)
  - Image push
  - kubectl delete
  - kubectl aaply

