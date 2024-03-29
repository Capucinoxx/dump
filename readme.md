## Développement en local

```sh
docker build -t wordle .
docker run -p 10436:8080 wordle
```

et se rendre sur http://localhost:10436/

## Test en local

Si vous souhaiter exécuter les tests en local sur votre machine:

```sh
docker build -t wordle-test . -f Dockerfile.test
docker run wordle-test
```

/!\ Dans la version initiale, cela donne une erreur sur `wordle_test.go:32`.
    Pour retirer ce comportement gênant, il faut corriger le test dans le code
    en golang.