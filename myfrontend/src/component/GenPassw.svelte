<script>
  import { Alert } from "sveltestrap";
  import { Button, Modal, ModalBody, ModalFooter, ModalHeader } from "sveltestrap";
  import axios from "axios";
  import Cookies from 'js-cookie';
  
//TODO: add alert when copying generated password
//TODO: if user is not loggeding don't show this page

  let open = false;
  const toggle = () => (open = !open);
  var copied = false;
  var islogged = true;
  let visible = true;
  let disable = true;
  var genPwData = { sliderValue:8, checkBoxVal:"" }
  var respData = { passwordVal:"" };
  let btncolor = "light";
  let cookieName = "session";
  let failcolor = "danger";
  let succolor = "success";
  let noType = false;

  function trackvalue() {
    var slider = document.getElementById("myRange");
    genPwData.sliderValue = slider.value;
  }

  function checkCookie() {
    var cookie = Cookies.get(cookieName);
    if (cookie != null) {
      islogged = true;
    } else {
      islogged = false;
    }
  }


  function getPwType(e) {
    if (open) {
      e.preventDefault();
      var checkBox = document.getElementById("checks");
      for (var i=0;i<checkBox.length;++i) {
        if (checkBox[i].checked) {
          genPwData.checkBoxVal = checkBox[i].value
        }
      }
    }
    open = !open
  }


  function getValueAndDisable() {
    if (genPwData.checkBoxVal !== "") {
      noType = false;
      var bodyFormData = new FormData();
      var resultField = document.getElementById("result");
      bodyFormData.set("type",genPwData.checkBoxVal);
      bodyFormData.set("length",genPwData.sliderValue);

      axios({
        method:'post',
        url:'http://localhost:8080/genpw',
        data:bodyFormData,
        responseType:'json',
        headers:{'Content-Type': 'multipart/form-data'}
      })
        .then((response) => {
          var resp = response.data;
          var r = JSON.stringify(resp);
          var obj = JSON.parse(r);
          respData.passwordVal = obj.Value;
          resultField.value = respData.passwordVal;
        })
        .catch((error) => {
          console.log(error)
        })
    } else {
      noType = true;
      visible=true
    }
    genPwData.checkBoxVal = "";
  }

  function CopyToClipboard() {
    var result = document.getElementById("result")
    if (result.value !== "") {
      result.disabled = false;result.select();
      result.setSelectionRange(0,99999);
      document.execCommand("copy");
      result.disabled=true;copied = true;
      visible=true;
    } else {
      copied=false;
    }
  }

</script>


<main>
  { #if copied }
    <Alert color={succolor} isOpen={visible} toggle={() => (visible = false,copied=false)}>
      Copied To clipboard
    </Alert>
  {/if}

  { #if noType }
    <Alert color={failcolor} isOpen={visible} toggle={() => (visible = false)}>
      No type has been specified
    </Alert>
  {/if}

  { #if !islogged }
    <Alert color={failcolor} isOpen={visible} toggle={() => (visible = false)}>
       You Need to login or register if you don't have an account <a href="/login">login</a> <a href="/register">register</a>
    </Alert>
  { :else } 
    <h4>Generate Your Password:</h4>
        <Modal id="modal" isOpen={open} {toggle}>
          <ModalHeader {toggle}>Choose a password type</ModalHeader>
          <ModalBody id="typepw">
            <form id="checks" on:submit={getPwType}>
              <div class="radioBtn">
                <div>
                  <input id="char" name="choice" type="radio" value=character> 
                  <label id="lchar">Characters</label>
                </div>
                <div>
                  <input id="int" name="choice" type="radio" value="integer"> 
                  <label id="lint">Integers</label>
                </div>
                <div>
                  <input id="special" name="choice" type="radio" value="special"> 
                  <label id="lsp">Special Characters </label>
                </div>
                <div>
                  <input id="mix" name="choice" type="radio" value="mix"> 
                  <label id="lmix">Mix Password </label>
                </div>
              </div>
              <Button id="ok" {btncolor}>OK</Button>
            </form>
          </ModalBody>
        </Modal>
        <div class="slidecontainer">
          <li id="pwL">Password Length</li>
          <input on:input={trackvalue} type="range" min="8" max="50" value="8" class="slider" id="myRange">
          <h5>{genPwData.sliderValue}</h5>
        </div>
        <div id="typeBtn">
          <Button  {btncolor} on:click={toggle}> Choose Password Type </Button>
        </div>
        <div>
          <Button name="btn" id="sub" {btncolor} on:click={getValueAndDisable}>Generate</Button>
        </div>
        <div>
          <input type="text" id="result" disabled={disable}>
          <Button id="copy" on:click={CopyToClipboard} {btncolor}>  📋  </Button>
        </div>
  {/if}
</main>


<style>

	.sub {
		background-color: #8cd7f4;
		padding: 16px 20px;
    padding-top: 18%;
		border: none;
		cursor: pointer;
		width: 50%;
		opacity: 0.9;
	}
	.sub:hover {
  	opacity:1;
	}
  
  #typeBtn {
    margin-top:2%;
    margin-bottom:2%;
  }


  h4{
    padding-top:2%;
    padding-bottom:1%;
  }

  h5 {
    padding-top:2%;
    padding-bottom:2%;
    text-align:center;
    color:#0ebf44;
  }

  #pwL, #typeli {
    color:#3574dc;
  }
  
  .radioBtn {
    padding-top:3%;
  }


  #char,#special,#int,#mix{
    float:left;
  }
  

  input[type="text"] {
    font-color:black;
    color:black;
    margin-top:2%;
    text-align:center;
    width:50%;
    padding-right:1%;
    padding-left:1%;
  }


  .slidecontainer {
    width: 78%; /* Width of the outside container */
    padding-left:25%;
    padding-top:2%;
  }



  /* The slider itself */
  .slider {
    -webkit-appearance: none;  
    appearance: none;
    width: 50%; 
    height: 25px; 
    background: #d3d3d3; 
    outline: none; 
    opacity: 0.7; 
    -webkit-transition: .2s; 
    transition: opacity .2s;
  }

  .slider:hover {
    opacity: 1; 
  }

  .slider::-webkit-slider-thumb {
    -webkit-appearance: none; 
    appearance: none;
    width: 25px; 
    height: 25px; 
    background: #4CAF50; 
    cursor: pointer; 
  }

  .slider::-moz-range-thumb {
    width: 25px; 
    height: 25px; 
    background: #4CAF50; 
    cursor: pointer; 
  }
</style>
