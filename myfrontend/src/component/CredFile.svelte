<script>
import axios from 'axios';
import Cookies from 'js-cookie';
import { Button } from 'sveltestrap';
import { Alert } from "sveltestrap";

let btncolor = "light";
let visible = true;
let supportedFiles = ".csv";
let succolor = "success";
let failcolor = "danger";
var lines;
var success;
var fail;
var isLogged;


function checkCookie() {
  var cookie = Cookies.get("session");
  if(cookie != null) {
    return true
  }
  return false
}

isLogged = checkCookie();

function UploadFile() {
  var f = document.getElementById("file").files;
  var bodyFormData = new FormData();
  bodyFormData.set("myfile",f[0]);
  if (f[0].type === "text/csv") {
    axios({
      method:"post",
      url:'http://localhost:8080/upload',
      data:bodyFormData,
      responseType:'json',
      withCredentials:true,
      headers:{ 'Content-Type': 'multipart/form-data' }
    })
      .then((response) => {
        var resp    = response.data;
        var r       = JSON.stringify(resp);
        var obj     = JSON.parse(r);
        console.log(obj)
        var isSave  = obj.Success;
        if (isSave) {
          success = true
        } else {
          fail = true
        } 
        lines = obj.Lines
      })
      .catch((error) => {
        console.log(error)
      });
  }
}

function submit(e) {
  e.preventDefault();
  UploadFile()
}

</script>


<main>
  { #if fail && lines.length != 0 }
    <Alert color={failcolor} isOpen={visible} toggle={() => (visible = false)}>
      You have upload some creds that already exist see line(s) => {lines.join()}
    </Alert>
  { /if }
  { #if success }
    <Alert color={succolor} isOpen={visible} toggle={() => (visible = false)}>
      Saved Success Fully
    </Alert>
  { /if }
  { #if !isLogged }
    <Alert color={failcolor} isOpen={visible} toggle={() => (visible = false)}>
      Loggin First <a href="/login">login</a>
    </Alert>
  { :else }
    <div id="forms">
      <form on:submit={submit} method="post">
        <input type="file" id="file">
        <Button name="btn" id="upload" {btncolor}>Upload</Button>
      </form>
    </div>
  { /if }
</main>

<style>

  #forms {
    padding-top:15%;
    padding-right:15%;
    padding-left:15%;
  }

</style>
