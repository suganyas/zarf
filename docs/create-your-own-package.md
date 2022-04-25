# Defining and Deploying Your Own Zarf Package

At this point you should be slightly familiar with Zarf.  If you're new to Zarf, you should start by reading the [What is Zarf](./what-is-zarf.md) page.


# Types of Zarf Packages
There are two types of Zarf packages, a `ZarfInitConfig` and a `ZarfPackageConfig`. The package type is defined by the `kind:` field in the zarf.yaml file. Zarf init configs were mentioned briefly in the [What is Zarf](./what-is-zarf.md) page as the package you use to  


The first type is a "base" package, which is a package that contains all of the files and directories that are required to run a Zarf application.  The second type is a "plugin" package, which is a package that contains additional functionality that can be added to an existing Zarf application.


## Creating a Zarf.yaml
Every Zarf package starts with a Zarf.yaml. This file defines the package's name, version, and other metadata.  It also defines the package's dependencies.  The Zarf.yaml file is located in the root of the package.

Most Zarf packages start the same as

```
This is a code block.

```
