window.onload = function (){
    fetch("/student/all")
    .then(response => response.text())
    .then(data => showAllStudents(data))
}

function newRow(table, student){
    var row = table.insertRow(table.length)
        
        var td =[];
        for(i = 0; i<table.rows[0].cells.length; i++){
            td[i] = row.insertCell(i);
        }

        // insert data in td
        td[0].innerHTML = student.stdid
        td[1].innerHTML = student.fname
        td[2].innerHTML = student.lname
        td[3].innerHTML = student.email
        td[4].innerHTML = '<input type="button" onclick="deleteStudent(this)" value="delete" id="button-1">'
        td[5].innerHTML = '<input type="button" onclick="updateStudent(this)" value="edit" id="button-2">'

}

function showAllStudents(data){
    const allStudents = JSON.parse(data)
    var table = document.getElementById("myTable")
    allStudents.forEach(stud => {
        newRow(table, stud)
    });

}


function showStudent(data){
    // console.log(data);
    // convert json string to js object 
    const student = JSON.parse(data)
    var table = document.getElementById("myTable")
    newRow(table, student)



        
}
// form reset 
function resetForm(){
    document.getElementById("sid").value = ""
    document.getElementById("fname").value = ""
    document.getElementById("lname").value = ""
    document.getElementById("email").value = ""
}

// helper function to get form data
function getFormData(){
    // create js object to store form data
      var formData = {
        stdid : parseInt(document.getElementById("sid").value),
        fname : document.getElementById("fname").value,
        lname : document.getElementById("lname").value,
        email : document.getElementById("email").value
    }
    return formData
}

function addStudent(){
    data = formData()
    // form validation

    var sid = data.stdid
    if(isNaN(sid)){
        alert("enter valid student ID")
        return
    }else if(data.email == ""){
        alert("Email cant be empty")
        return
    }else if(data.fname == ""){
        alert("First name cant be empty")
        return
    }
        
    // call POST API
    fetch('/student/add', {
        method: "POST",
        body: JSON.stringify(data),
        headers: {"Content-type": "application/json; charset=UTF-8"}
    }).then(response1 => {
        // check response from the fetch 
        if(response1.ok){
            fetch("/student/"+data.stdid)
            .then(response2 => response2.text())
            .then(data => showStudent(data))
        }else{
            throw new Error(response1.statusText)
        }
    }).catch(e => {
        if (e.message == 401){
            alert("User not not loged in")
            window.open("index.html", "_self")
        } else if (e.message = 400){
            alert("Bad Request")
        }else{
            alert("Internal server error")
        }
    } );
    resetForm()
}

function updateAPIRequest(oldSid){
    // get the updated data from user 
    newData = getFormData()
    // call update api
    fetch("/student/"+oldSid, {
        method: "PUT",
        body: JSON.stringify(newData),
        headers: {"Content-type": "application/json; charset=UTF-8"}
    }).then(res => {
        if(res.ok){
            selectedRow.cells[0].innerHTML = newData.stdid
            selectedRow.cells[1].innerHTML = newData.fname
            selectedRow.cells[2].innerHTML = newData.lname
            selectedRow.cells[3].innerHTML = newData.email

            // change button value to initiual state
            var btn = document.getElementById("button-add")
            btn.innerHTML = "Add"
            btn.setAttribute("onclick", "addStudent()")

            selectedRow = null
            resetForm()
        }else{
            alert("server: update request error")
        }
    })

}

var selectedRow = null

function updateStudent(input){
    // get the selected row
    selectedRow = input.parentElement.parentElement
    document.getElementById("sid").value = selectedRow.cells[0].innerHTML
    document.getElementById("fname").value = selectedRow.cells[1].innerHTML
    document.getElementById("lname").value = selectedRow.cells[2].innerHTML
    document.getElementById("email").value = selectedRow.cells[3].innerHTML

    sid = selectedRow.cells[0].innerHTML

    // change button value to update
    var btn = document.getElementById("button-add")
    btn.innerHTML = "Update"
    btn.setAttribute("onclick", "updateAPIRequest(sid)")
}

function deleteStudent(input){
    if(confirm("Are you sure to delete this?")){
        selectedRow = input.parentElement.parentElement
        sid = selectedRow.cells[0].innerHTML
        fetch("/student/"+sid, {
            method: "DELETE"
        })
        .then(res => {
            if(res.ok){
                var rowIndex = selectedRow.rowIndex
                document.getElementById("myTable").deleteRow(rowIndex)
                alert("student deleted successfully")
                selectedRow = null
            }else{
                alert("server: delete request error")
            }
        })
    }
}