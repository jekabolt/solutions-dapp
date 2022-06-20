class HttpAPI {
    constructor(singleAirdropForm,
        multipleAirdropForm,
        listAllForm,
        listFailedForm,
        listPendingForm,
    ) {
        this.singleAirdropForm = singleAirdropForm;
        this.multipleAirdropForm = multipleAirdropForm;

        this.listAllForm = listAllForm;
        this.listFailedForm = listFailedForm;
        this.listPendingForm = listPendingForm;

        this.submitSingleAirdrop()
        this.submitMultipleAirdrops()
        this.listAllAirdrops()
        this.listFailedAirdrops()
        this.listPendingAirdrops()


        this.post = function(url, data) {
            return fetch(url, {
                method: "POST",
                headers: {
                    'Content-Type': 'application/json',
                    'Grpc-Metadata-Authorization': `Bearer ${localStorage.getItem("token")}`
                },
                body: JSON.stringify(data)
            });
        }
        this.get = function(url) {
            return fetch(url, {
                method: "GET",
                headers: {
                    'Accept': 'application/json',
                    'Grpc-Metadata-Authorization': `Bearer ${localStorage.getItem("token")}`
                }
            });
        }
        this.delete = function(url) {
            return fetch(url, {
                method: "DELETE",
                headers: {
                    'Content-Type': 'application/json',
                    'Grpc-Metadata-Authorization': `Bearer ${localStorage.getItem("token")}`
                }
            });
        }
        this.put = function(url, data) {
            return fetch(url, {
                method: "PUT",
                headers: {
                    'Content-Type': 'application/json',
                    'Grpc-Metadata-Authorization': `Bearer ${localStorage.getItem("token")}`
                },
                body: JSON.stringify(data)
            });
        }

        this.handleFormButton = function(form) {
            var elements = form.getElementsByTagName("input");
            for (var ii = 0; ii < elements.length; ii++) {
                if (elements[ii].type == "text") {
                    elements[ii].value = "";
                }
            }

            let button = form.getElementsByClassName("button")[0]
            button.innerHTML = "ok";
            button.disabled = true;
            button.style.color = "gray";
            setTimeout(function() {
                button.innerHTML = "submit";
                button.disabled = false
                button.style.color = "white";
            }, 2000)
        }
        this.download = function(content, filename, contentType) {
            var blob = new Blob([content], { type: contentType });
            var url = URL.createObjectURL(blob);

            var pom = document.createElement('a');
            pom.href = url;
            pom.setAttribute('download', filename);
            pom.click();

        }

        this.CSVToJSON = csv => {
            const removeEmptyLines = str => str.split(/\r?\n/).filter(line => line.trim() !== '').join('\n');
            const lines = removeEmptyLines(csv).split('\n');
            const keys = lines[0].split(',');
            return lines.slice(1).map(line => {
                return line.split(',').reduce((acc, cur, i) => {
                    const toAdd = {};
                    if (cur == 'true') {
                        toAdd[keys[i]] = true;
                    } else if (cur == 'false') {
                        toAdd[keys[i]] = false;
                    } else if (cur == '') {

                    } else {
                        toAdd[keys[i]] = cur;
                    }
                    return {...acc, ...toAdd };
                }, {});
            });
        }
        this.JSONToCSV = (jsonarray) => {
            var csvrecord = Object.keys(jsonarray[0]).join(',') + '\n';
            jsonarray.forEach(function(jsonrecord) {
                csvrecord += Object.values(jsonrecord).join(',') + '\n';
            });
            return csvrecord
        };

        this.logout = function() {
            localStorage.removeItem("token");
            document.location.reload(true)
        }

        this.handleGET = function(resp, form) {
            let self = this;
            resp.then(function(response) {
                switch (response.status) {
                    case 200:
                        response.json().then(function(data) {
                            let today = new Date().toISOString()
                            self.download(self.JSONToCSV(data.airdrops), today + ".csv", 'text/csv;charset=utf-8;')
                            self.handleFormButton(form);
                        });
                        break;
                    case 401:
                        alert("token expired logout")
                        self.logout()
                        break;

                    default:
                        alert("server error")
                }

            })
        }
        this.handlePOST = function(resp, form) {
            let self = this;
            resp.then(function(response) {
                switch (response.status) {
                    case 200:
                        response.json().then(function(data) {
                            console.log(data);
                            self.handleFormButton(form);
                        });
                        break;
                    case 401:
                        alert("token expired logout")
                        self.logout()
                        break;

                    default:
                        alert("server error or bad input")
                }

            })
        }

    }
    submitSingleAirdrop() {
        let self = this;
        this.singleAirdropForm.addEventListener("submit", (e) => {
            e.preventDefault();
            var error = 0;

            let resp = self.post("/api/airdrop/bulk-submit", {
                airdrops: [{
                    amount: parseInt(document.getElementById("amount").value),
                    currency: document.getElementById("currency").value,
                    email: document.getElementById("email").value,
                    purpose: document.getElementById("purpose").value,
                    supernode: document.getElementById("supernodeId").value,
                }],
                authToken: localStorage.getItem("token"),
            })
            self.handlePOST(resp, self.singleAirdropForm)
        });
    }

    submitMultipleAirdrops() {
        let self = this;
        this.multipleAirdropForm.addEventListener("submit", (e) => {
            e.preventDefault();
            var error = 0;

            let self = this;
            const formFile = self.multipleAirdropForm.querySelector('#input_file');
            if (!window.File || !window.FileReader || !window.FileList || !window.Blob) {
                console.log('The File APIs are not fully supported in this browser.');
                return;
            }

            if (!formFile.files) {
                console.log("This browser doesn't seem to support the `files` property of file inputs.");
            } else if (!formFile.files[0]) {
                console.log("No file selected.");
            } else {
                let file = formFile.files[0];
                let fr = new FileReader();
                fr.onload = receivedText;
                fr.readAsText(file);

                function receivedText() {
                    let resp = self.post("/api/airdrop/bulk-submit", {
                        airdrops: self.CSVToJSON(fr.result),
                        authToken: localStorage.getItem("token"),
                    })
                    self.handlePOST(resp, self.multipleAirdropForm)
                }
            }

        });
    }

    listAllAirdrops() {
        let self = this;
        this.listAllForm.addEventListener("submit", (e) => {
            e.preventDefault();

            let resp = self.get("/api/airdrop/list")
            self.handleGET(resp, self.listAllForm)

        });
    }

    listFailedAirdrops() {
        let self = this;
        this.listFailedForm.addEventListener("submit", (e) => {
            e.preventDefault();

            let resp = self.get("/api/airdrop/list-failed")
            self.handleGET(resp, self.listFailedForm)
        });
    }

    listPendingAirdrops() {
        let self = this;
        this.listPendingForm.addEventListener("submit", (e) => {
            e.preventDefault();

            let resp = self.get("/api/airdrop/list-pending")
            self.handleGET(resp, self.listPendingForm)
        });
    }

}


const singleAirdropForm = document.querySelector(".single-airdrop");
const multipleAirdrop = document.querySelector(".multiple-airdrop");

const listAllForm = document.querySelector(".list-all-airdrops");
const listFailedForm = document.querySelector(".list-failed-airdrops");
const listPendingForm = document.querySelector(".list-pending-airdrops");

if (singleAirdropForm && multipleAirdrop && listAllForm && listFailedForm && listPendingForm) {
    const api = new HttpAPI(
        singleAirdropForm,
        multipleAirdrop,
        listAllForm,
        listFailedForm,
        listPendingForm,
    );
}