---
layout: default
title: Expressions
nav_order: 4
---

### Expressions

You can use expressions to programmatically set environment variables in workflow files and access contexts. An expression can be any combination of literal values, references to a context, or functions. You can combine literals, context references, and functions using operators. For more information about contexts, see "Contexts."

#### Literals

You can use boolean, null, number, or string data types.

<details>
  <summary>Example of literals</summary>

```yml
env:
  myNull: ${{ null }}
  myBoolean: ${{ false }}
  myIntegerNumber: ${{ 711 }}
  myFloatNumber: ${{ -9.2 }}
  myHexNumber: ${{ 0xff }}
  myExponentialNumber: ${{ -2.99e-2 }}
  myString: Mona the Octocat
  myStringInBraces: ${{ 'It''s open source!' }}
```
</details>

#### Operators

<details>
  <summary>Example of operators</summary>

```
Operator	Description
( )	Logical grouping
[ ]	Index
.	Property de-reference
!	Not
<	Less than
<=	Less than or equal
>	Greater than
>=	Greater than or equal
==	Equal
!=	Not equal
&&	And
||	Or
```
</details>

> [!TIP]
> You can use a ternary operator `condition ? true : false` as `${{ condition && true || false }}`.

[Expressions](https://docs.github.com/en/actions/writing-workflows/choosing-what-your-workflow-does/expressions)

#### Functions

You can use functions to transform data or to perform operations.

* [contains](https://docs.github.com/en/actions/writing-workflows/choosing-what-your-workflow-does/expressions#contains)
* [startswith](https://docs.github.com/en/actions/writing-workflows/choosing-what-your-workflow-does/expressions#startswith)
* [endsWith](https://docs.github.com/en/actions/writing-workflows/choosing-what-your-workflow-does/expressions#endswith)
* [format](https://docs.github.com/en/actions/writing-workflows/choosing-what-your-workflow-does/expressions#format)
* [join](https://docs.github.com/en/actions/writing-workflows/choosing-what-your-workflow-does/expressions#join)
* [toJson](https://docs.github.com/en/actions/writing-workflows/choosing-what-your-workflow-does/expressions#tojson)
* [fromJson](https://docs.github.com/en/actions/writing-workflows/choosing-what-your-workflow-does/expressions#fromjson)
* [hashFiles](https://docs.github.com/en/actions/writing-workflows/choosing-what-your-workflow-does/expressions#hashfiles)

#### Status Check functions

* [success](https://docs.github.com/en/actions/writing-workflows/choosing-what-your-workflow-does/expressions#success)
* [always](https://docs.github.com/en/actions/writing-workflows/choosing-what-your-workflow-does/expressions#always)
* [cancelled](https://docs.github.com/en/actions/writing-workflows/choosing-what-your-workflow-does/expressions#cancelled)
* [failure](https://docs.github.com/en/actions/writing-workflows/choosing-what-your-workflow-does/expressions#failure)
