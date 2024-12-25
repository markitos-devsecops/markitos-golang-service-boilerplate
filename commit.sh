#!/bin/bash

# Solicita el mensaje de commit en español
read -p "Introduce el mensaje de commit en español: " commit_message

# Añade un guion delante del mensaje
commit_message="- $commit_message"

# Añade todos los cambios, hace el commit y el push
git add .
git commit -m "$commit_message"
git push