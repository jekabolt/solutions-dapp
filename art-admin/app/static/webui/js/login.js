class Login {
    constructor(form, fields) {
        this.form = form;
        this.fields = fields;
        this.username;
        this.password;
        this.validateOnSubmit();
        this.post = function(url, data) {
            return fetch(url, { method: "POST", headers: { 'Content-Type': 'application/json' }, body: JSON.stringify(data) });
        }
        this.logout();
    }
    logout() {
        document.getElementById("logout").addEventListener("click", () => {
            localStorage.removeItem("token");
            document.location.reload(true)
        });
    }
    validateOnSubmit() {
        let self = this;
        this.form.addEventListener("submit", (e) => {
            e.preventDefault();
            var error = 0;

            self.fields.forEach((field) => {
                const input = document.querySelector(`#${field}`);
                if (self.validateFields(input) == false) {
                    error++;
                }
            });

            self.username = document.getElementById("username").value
            self.password = document.getElementById("password").value

            self.post("/api/auth/login", {
                password: self.password,
                username: self.username,
            }).then(function(response) {
                if (response.status == 200) {
                    response.json().then(function(data) {
                        localStorage.setItem("token", data.authToken);
                        console.log(data);
                    });
                } else {
                    document.getElementById("password").classList.add("input-error");
                    document.getElementById("password").parentElement.querySelector(".error-message").innerText = "Incorrect username or password or username";
                }
            }).then(function() {
                if (error == 0) {
                    window.location.reload();
                }
            })
        });
    }
    validateFields(field) {
        if (field.value.trim() === "") {
            this.setStatus(
                field,
                `${field.previousElementSibling.innerText} cannot be blank`,
                "error"
            );
            return false;
        } else {
            if (field.type == "password") {
                if (field.value.length < 8) {
                    this.setStatus(
                        field,
                        `${field.previousElementSibling.innerText} must be at least 8 characters`,
                        "error"
                    );
                    return false;
                } else {
                    this.setStatus(field, null, "success");
                    return true;
                }
            } else {
                this.setStatus(field, null, "success");
                return true;
            }
        }
    }
    setStatus(field, message, status) {
        const errorMessage = field.parentElement.querySelector(".error-message");

        if (status == "success") {
            if (errorMessage) {
                errorMessage.innerText = "";
            }
            field.classList.remove("input-error");
        }
        if (status == "error") {
            errorMessage.innerText = message;
            field.classList.add("input-error");
        }
    }
}

if (localStorage.getItem("token")) {
    document.getElementById("login").style.display = "none";
} else {
    document.getElementById("dashboard").style.display = "none";
}

const form = document.querySelector(".loginForm");
if (form) {

    const fields = ["username", "password"];

    const validator = new Login(form, fields);
}