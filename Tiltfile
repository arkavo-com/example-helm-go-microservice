# Tiltfile for development
# https://docs.tilt.dev/api.html

docker_build('microservice', '.', target='runner')

# helm
k8s_yaml(helm('helm', 'ex'))
