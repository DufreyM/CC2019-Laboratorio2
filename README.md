Video: https://youtu.be/2477xUYGxmc 

El algoritmo de Shunting Yard, es una técnica para convertir expresiones en notación infix (como a|(b*c)) a notación postfix o notación polaca inversa (como abc*|). Esta transformación es útil para procesar expresiones con operadores y paréntesis de manera estructurada, eliminando la necesidad de gestionar precedencia y paréntesis durante la evaluación.

En el contexto de expresiones regulares, este algoritmo permite convertir una expresión con operadores como |, *, +, ?, ^, y paréntesis, en una forma lineal que facilita su análisis posterior (por ejemplo, para construir autómatas finitos).

El algoritmo utiliza dos estructuras principales:

Una pila de operadores
Una lista de salida (postfix)