Video: https://youtu.be/2477xUYGxmc 

El algoritmo de Shunting Yard, es una técnica para convertir expresiones en notación infix (como a|(b*c)) a notación postfix o notación polaca inversa (como abc*|). Esta transformación es útil para procesar expresiones con operadores y paréntesis de manera estructurada, eliminando la necesidad de gestionar precedencia y paréntesis durante la evaluación.

En el contexto de expresiones regulares, este algoritmo permite convertir una expresión con operadores como |, *, +, ?, ^, y paréntesis, en una forma lineal que facilita su análisis posterior (por ejemplo, para construir autómatas finitos).

El algoritmo utiliza dos estructuras principales:

Una pila de operadores
Una lista de salida (postfix)

### Ejercicio No. 1 - Conversión de Expresiones Regulares a Autómatas Finitos

El archivo `Ejercicio1.pdf` contiene las soluciones detalladas de los siguientes ejercicios, donde se convirtieron expresiones regulares a Autómatas Finitos Deterministas (AFD) mediante el proceso de construcción de AFN con Thompson, tablas de transición y conversión a AFD:

**Expresiones regulares resueltas:**
1. `(𝑎|𝑡)𝑐`
2. `(𝑎|𝑏) ∗`
3. `(𝑎 ∗ |𝑏 ∗) ∗`
4. `((𝜀|𝑎)|𝑏 ∗) ∗`
5. `(𝑎|𝑏) ∗ 𝑎𝑏𝑏(𝑎|𝑏) ∗`
6. `0? (1? )? 0 ∗`
7. `𝑖𝑓\([𝑎𝑒] +\)\{[𝑒𝑖] +\}(\𝑛(𝑒𝑙𝑠𝑒\{[𝑗𝑙] +\}))?`
8. `[𝑎𝑒03] + @[𝑎𝑒03]+. (𝑐𝑜𝑚|𝑛𝑒𝑡|𝑜𝑟𝑔)(. (𝑔𝑡|𝑐𝑟|𝑐𝑜))?`

**Procedimiento incluido en el PDF:**
- Construcción de AFN usando el algoritmo de Thompson.
- Generación de tablas de transición.
- Conversión de AFN a AFD.
