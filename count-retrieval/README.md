# Frontend

## Info
The frontend uses [Angular](https://angular.io/) for rendering and routing. The db facade lives `/src/app/services`.

## Local Development
**You will need both `git` and `npm` installed to run this project.**

First run  `npm install -g @angular/cli` and then  run`npm install` in the `frontend` directory. After that is finished, run `ng serve` to start a local development server. Visit 
`localhost:8080` to see the page.

### Formatting
Please run:
`prettier --write "**/*.ts"`
and
`prettier --write "**/*.html"`
from the root directory before making a pull request

## Hints
If you get an error regarding missing dependcies, please run `npm install -g @angular/cli@1.7.1`
