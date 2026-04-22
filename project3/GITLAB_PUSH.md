# Подключение проекта к GitLab и запуск pipeline

## Вариант 1. Создать пустой проект через UI GitLab

```bash
git init
git branch -M main
git add .
git commit -m "feat: initial pr3 project name_nsa_13"
git remote add origin <URL_ВАШЕГО_GITLAB_REPO>
git push -u origin main
```

## Вариант 2. Создать проект сразу через git push

```bash
git init
git branch -M main
git add .
git commit -m "feat: initial pr3 project name_nsa_13"
git push --set-upstream https://gitlab.com/<namespace>/<project>.git main
```

## После push

Откройте Build -> Pipelines.
Нужен зелёный статус для jobs: lint, sast, test, build.

## Если pipeline не стартует

- проверьте, что CI/CD включен
- проверьте, что для проекта доступен runner
- проверьте, что `.gitlab-ci.yml` лежит в корне проекта

## Локальная проверка

```bash
chmod +x local_ci.sh
./local_ci.sh
```
