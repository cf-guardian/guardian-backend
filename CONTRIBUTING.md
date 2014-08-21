# Contribution Guidelines

The Cloud Foundry team uses GitHub and accepts contributions via
[pull request](https://help.github.com/articles/using-pull-requests).

## Contributor License Agreement

Follow these steps to make a contribution to any of our open source repositories:

1. Ensure that you have completed our CLA Agreement for
  [individuals](http://www.cloudfoundry.org/individualcontribution.pdf) or
  [corporations](http://www.cloudfoundry.org/corpcontribution.pdf).

1. Set your name and email (these should match the information on your submitted CLA)

    ```shell
git config --global user.name "Firstname Lastname"
git config --global user.email "your_email@example.com"
    ```

## General Workflow

1. Fork the repository,
1. create a branch (`git checkout -b my_feature`),
1. [write your changes and run the tests](DEVELOPMENT.md#testing),
1. commit changes onto your branch, and
1. push to your fork (`git push origin my_feature`) and submit a pull request.

Small pull requests with an obvious, single purpose are preferred.

Your pull request is also more likely to be accepted if it:

* is small;
* consists of a single commit;
* includes tests;
* conforms to standard Go formatting conventions (`go fmt`); and
* contains a message explaining the intent of your change.
