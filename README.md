# ws_neologd_normalize - MySQL UDF for neologd normalize.

### about

This is MySQL User Defined Function written by cgo.  
Normalize string following neologd.

(neologd normalize algorithm)  

https://github.com/neologd/mecab-ipadic-neologd/wiki/Regexp.ja

In addition, this udf remove `\n`, `\r`, `\t` and `\v`.  

### how to install

    $ ./build.sh

(notice)

* require root privilege

### example

    MariaDB [(none)]> select mecab_normalize('　　　ＰＲＭＬ　　副　読　本　　　');
    +------------------------------------------------------------------------+
    | mecab_normalize('　　　ＰＲＭＬ　　副　読　本　　　')                  |
    +------------------------------------------------------------------------+
    | PRML副読本                                                             |
    +------------------------------------------------------------------------+
