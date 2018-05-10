# Gmail

## Development


Into ```go-commands```

    export GOPATH=`pwd`

    
and then...

    
    go get gopkg.in/kyokomi/emoji.v1



... export AWS_KEY and AWS_SECRET with full access on S3

    export GMAIL_EMAIL="YOUR_EMAIL_HERE"
    export GMAIL_PWD="YOUR_PASSWORD_HERE"


## Running


__Sending an email__

... to create a bucket, go into ```gmail``` after the __Development__ steps and run:

    go run main.go send TO_EMAIL, SUBJECT, HTML_FILE


## Generating 

    go build main.go
    mv main gmail


## HTML File Example


```
<!DOCTYPE HTML PULBLIC "-//W3C//DTD HTML 4.01 Transitional//EN">
<html>
<head>
<meta http-equiv="content-type" content="text/html"; charset="ISO-8859-1">
</head>
<body>
<b>Howdy!</b><br>
&nbsp;<br>
Eu sou o robo do pessoal de tecnologia da R/GA.<br>
Este email é para te avisar que o ambiente: <a href="https://rga.com">https://rga.com</a> 
está atualizado e você já pode dar uma olhada em como está ficando o progresso. <br>
<div class="moz-signature"><i><br>
<br>
Regards<br>
<br>
R/GA SP Tech 
<i></div>
</body>
</html>
```
