import { useState } from 'react';
import {CreateUserRequest} from '../lib/hello_pb';
import {HelloServiceClient} from '../lib/HelloServiceClientPb';
import { BASEURL } from './const/form_data';
import { BUCKET } from './const/form_data';
import AWS from 'aws-sdk';

export const UserRegisterForm = () => {
  const [name, setName] = useState('');
  const [score, setScore] = useState('');
  const [profileImage, setProfileImage] = useState('');
  const [imageFile, setImage] = useState('');
  const [imageUploadUrl, setImageUploadUrl] = useState('');

  const s3  = new AWS.S3({
    accessKeyId: 'minio' ,
    secretAccessKey: 'minio123' ,
    endpoint: 'http://127.0.0.1:9000' ,
    s3ForcePathStyle: true, // needed with minio?
    signatureVersion: 'v4'
  });

  //画像を選択した際の処理
  const processImage = (event) => {
    const File = event.target.files[0];
    setImage(File);
    const imageUrl = URL.createObjectURL(File);
    setProfileImage(imageUrl);
    const uploadURL =  BASEURL + File.name;
    setImageUploadUrl(uploadURL);
  }

  //画像アップロードメソッド
  const uploadImg = (file) => {
    var fileName = file.name;
    var params = {
        Bucket: BUCKET,
        Key: fileName,
        ContentType: file.type,
        Body: file
      }
    s3.putObject(params, function(err, data) {
      if (err)
        console.log(err)
      else
        console.log("Successfully uploaded data to testbucket/testobject");
    });
  }

  //フォームsubmit時の処理
  const submitUserRegisterForm = async(event) => {
    event.preventDefault();
    //画像アップロード
    uploadImg(imageFile);
    const request = new CreateUserRequest();
    request.setName(name);
    request.setScore(score);
    //DBには画像のURLを保存
    request.setPhotourl(imageUploadUrl);
    const client = new HelloServiceClient("http://localhost:8080");
    const response = await client.createUser(request, {});
  }
  return (
    <div className="user-form-container">
      <h4>user-register-form</h4>
      <form onSubmit={submitUserRegisterForm}>
        <div>
          <p>名前</p>
          <input
            type ="text"
            onChange={(e)=>setName(e.target.value)}
          />
        </div>
        <div>
          <p>スコア</p>
          <input
            type ="text"
            onChange={(e)=>setScore(e.target.value)}
          />
        </div>
        <div>
          <p>プロフィール画像アップロード</p>
          <input type="file" accept="image/*" onChange={processImage}></input>
        </div>
        <div className="profile-image-sec">
          <img src={profileImage} className="profile-image"></img>
        </div>
        <div>
          <button>送信</button>
        </div>
      </form>
      <style jsx>{`
      .user-form-container {
        margin:10px;
      }
      .profile-image-sec{
        margin:10px;
      }
      .profile-image {
        width: 100px;
        height: 100px;
      }
      `}
      </style>
    </div>
  )
}