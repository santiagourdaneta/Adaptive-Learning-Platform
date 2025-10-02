# 🧠 Adaptive Learning Platform

![Go](https://img.shields.io/badge/go-%2300ADD8.svg?style=for-the-badge&logo=go&logoColor=white)
![Python](https://img.shields.io/badge/python-3670A0?style=for-the-badge&logo=python&logoColor=ffdd54)
![Rust](https://img.shields.io/badge/rust-%23000000.svg?style=for-the-badge&logo=rust&logoColor=white)
![Gin Gonic](https://img.shields.io/badge/gin--gonic-008080?style=for-the-badge&logo=gin&logoColor=white)
![SQLite](https://img.shields.io/badge/sqlite-%2307405E.svg?style=for-the-badge&logo=sqlite&logoColor=white)
![HTML5](https://img.shields.io/badge/html5-%23E34F26.svg?style=for-the-badge&logo=html5&logoColor=white)
![CSS3](https://img.shields.io/badge/css3-%231572B6.svg?style=for-the-badge&logo=css3&logoColor=white)
![JavaScript](https://img.shields.io/badge/javascript-%23323330.svg?style=for-the-badge&logo=javascript&logoColor=%23F7DF1E)

## 📖 Descripción del Proyecto

Esta es una **plataforma de aprendizaje adaptativo** que utiliza IA para ofrecer rutas de estudio personalizadas. El proyecto se basa en una arquitectura de microservicios, combinando la velocidad y concurrencia de Go y Rust con la capacidad de procesamiento de IA de Python.

---

## 🚀 Arquitectura

El sistema se compone de tres servicios principales:

-   **Go API**: El backend principal, que gestiona las rutas, la lógica de negocio y la conexión con la base de datos.
-   **Python AI**: Un servicio de Flask encargado de generar rutas de aprendizaje personalizadas utilizando modelos de Machine Learning.
-   **Rust Engine**: Un motor de tiempo real que maneja la comunicación WebSocket para notificaciones instantáneas y características interactivas.

---

## 🛠️ Tecnologías Utilizadas

-   **Backend**: Go (Gin-Gonic), `database/sql` con SQLite.
-   **Inteligencia Artificial**: Python (Flask, scikit-learn).
-   **Comunicación en Tiempo Real**: Rust (Actix-Web, Actix).
-   **Frontend**: HTML, CSS (Pico.css), JavaScript.

---

## 📦 Cómo Ejecutar el Proyecto

1.  **Clonar el repositorio:**
    ```bash
    git clone https://github.com/santiagourdaneta/Adaptive-Learning-Platform
    cd Adaptive-Learning-Platform
    ```

2.  **Iniciar el servicio de Go:**
    ```bash
    cd go-api
    go mod tidy
    go run .
    ```

3.  **Iniciar el servicio de Python:**
    ```bash
    cd ../python-ai
    pip install -r requirements.txt
    python main.py
    ```

4.  **Iniciar el servicio de Rust:**
    ```bash
    cd ../rust-engine
    cargo run
    ```

5.  **Abrir el navegador:**
    Con todos los servicios en marcha, accede a la plataforma en tu navegador:
    ```
    http://localhost:8080/
    ```
---

## 🔖 Etiquetas y Hashtags

**Labels (etiquetas):**
- `go`
- `python`
- `rust`
- `microservices`
- `ai`
- `machine-learning`
- `education`
- `web-development`
- `api`
- `real-time`
- `websocket`

**Tags:**
- `adaptive-learning`
- `edtech`
- `go-api`
- `rust-websocket`
- `python-ai`
- `full-stack`
- `sqlite`

**Hashtags:**
- `#AdaptiveLearning`
- `#AIinEducation`
- `#GoLang`
- `#Rust`
- `#Python`
- `#Microservices`
- `#EdTech`
- `#FullStack`