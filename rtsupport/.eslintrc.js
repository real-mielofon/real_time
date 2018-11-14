export default {
    "rules": {
        "indent": [2, 4],
        "quotes": [2, "single"],
        "linebreak-style": [2, "windows"],
        "semi": [2, "always"]
    },
    "env": {
        "es6": true,
        "browser": true
    },
    "extends": "eslint:recommended",
    //        "extends": "standard"
    "ecmaFeatures": {
        "jsx": true,
        "experimentalObjectRestSpread": true
    },
    "plugins": [
        "react"
    ],
    "parser": "babel-eslint",
    "parserOptions": {
        "sourceType": "module",
        "allowImportExportEverywhere": false,
        "codeFrame": true
    }

};