Guten AUR, a Arch Linux AUR development helper
===============================================

[Homepage](https://github.com/gutenye/gutaur) |
[Documentation](https://github.com/gutenye/gutaur/wiki) |
[Issue Tracker](https://github.com/gutenye/gutaur/issues) |
[MIT License](http://choosealicense.com/licenses/mit) |
[by Guten](http://guten.me) |
[Gratipay](https://gratipay.com/gutenye) |
[Bountysource](https://www.bountysource.com/teams/gutenye)

|                |                                                            |
|----------------|------------------------------------------------------------|
| **Install**    |                                                            |
| ArchLinux      | `$ pacaur -S gutaur`                                       |

This is a part of [gutpackaging](https://github.com/gutenye/gutpackaging) project.

Getting started
---------------

**Settings**

```
AUR_CONFIG=$HOME/.gutaurrc
AUR_USERNAME=x
AUR_PASSWORD=y
```

You can set them

* Via Environment Variable
* Via ~/.gutaurrc

**USAGE**

```
$ create a PKGBUILD with _watch variable

  _watch=("https://github.com/gutenye/gutbackup/tags" "v([-.\d]+).tar.gz")

$ gutaur build
> Build a package.

$ gutaur source
> Generate a source-only tarball.

$ gutaur upload
> Upload source-only tarball to AUR website.

$ gutaur publish
> Invoke source and upload.

$ gutaur publish 1.0.3
> Download upstream 1.0.3 version and publish it.

$ gutaur autopublish
> Check upstream version, if found a new one, publish it.
```

Development
===========

Contributing
-------------

* Submit any bugs/features/ideas to github issue tracker.

Thanks to [all contributors](https://github.com/gutenye/gutaur/contributors).

Copyright
---------

The MIT License (MIT)

Copyright (c) 2015 Guten Ye

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.
