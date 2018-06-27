var container = document.getElementById("main");

function initRoutes() {
    page("/",initPage);
    page("/users", getUsers);
    page("/users/new", newUser);
}

function initPage() {
    container.innerHTML = `
        <h1>Bienvenido a la aplicación de gestión de usuarios</h1>
    `
}

function getUsers(ctx) {

    let header = new Headers({
        'Access-Control-Allow-Origin':'*',
        'Content-Type': 'application/json'
    });

    let sentData={
        method: 'GET',
        mode: 'cors',
        header: header,
        credential: "omit"
    };

    let url = "http://127.0.0.1:8080/api/users";

    fetch(url, sentData)
      .then( (res) => res.json())
      .then( (res) => showContent(res.content) )
}

function showContent(content) {
    container.innerHTML = "<h1> Todos los usuarios registrados</h1>"

    content.forEach( (e) => {
        container.innerHTML += `
            <div>
                <h2>${e.nombre}</h2>
                <p><a href="#">${e.email}</a></p>
            </div>
        `
    })
}

function newUser(ctx) {


    container.innerHTML = `
        <h1>Crear usuario</h1>
        <form> 
            <div>
                <label>
                    Nombre: <input type="text" id="nombre" placeholder="Ingresar nombre de usuario" required>
                </label>    
            </div>
            <div>
                <label>
                    Email: <input type="email" id="email" placeholder="Ingresar email de usuario" required>
                </label>
            </div>
            <div>
                <label>
                    Password: <input type="password" id="password" placeholder="Ingresar password" required>
                </label>
            </div>
            <div>
                <label>
                    Telefono: <input type="number" id="telefono" placeholder="Ingresar telefono" required>
                </label>
            </div>
            <div>
                <label>
                    Pais: <input type="text" id="pais" placeholder="Ingresar pais de residencia">
                </label>
            </div>
            <div>
                <label>
                    Ciudad: <input type="text" id="ciudad" placeholder="Ingresar ciudad de residencia">
                </label>
            </div>
            <div>
                <label>
                    Dirección: <input type="text" id="direccion" placeholder="Ingresar dirección de residencia">
                </label>
            </div>
            <div>
                <input type="submit" value="Registrar usuario" id="btn">
            </div>
            
        </form>
    `;

    let btn = document.getElementById("btn")
    btn.addEventListener("click", (ev) => {
        ev.preventDefault()

        let data = {
            nombre: document.getElementById("nombre").value,
            email: document.getElementById("email").value,
            password: document.getElementById("password").value,
            telefono: document.getElementById("telefono").value,
            pais: document.getElementById("pais").value,
            ciudad: document.getElementById("ciudad").value,
            direccion: document.getElementById("direccion").value
        }


        let url = "http://127.0.0.1:8080/api/users"

        fetch(url, {
            method: "post",
            body: JSON.stringify(data),
            headers: {
                "Content-Type": "application/json"
            },
            credentials: "omit"
        })
            .then( (res) => res.json())
            .catch( (err) => console.error(err) )
            .then( (res) => {
                let pResponse = document.getElementById("message")
                pResponse.textContent = `${res.message}`
            });
    })

}

