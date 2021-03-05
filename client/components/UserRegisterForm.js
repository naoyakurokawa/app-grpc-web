import { useState } from 'react';
import {CreateUserRequest} from '../lib/hello_pb';
import {HelloServiceClient} from '../lib/HelloServiceClientPb';
import { BASEURL } from './const/form_data';
import { BUCKET } from './const/form_data';
import axios from 'axios';
import * as minio from "minio";
// import { useForm } from "react-hook-form";

export const UserRegisterForm = () => {
  // console.log('aaaaaaaaaaa',BASEURL)
  const [name, setName] = useState('');
  const [score, setScore] = useState('');
  const [profileImage, setProfileImage] = useState('');
  const [imageFile, setImage] = useState('');
  const [imageUploadUrl, setImageUploadUrl] = useState('');

  //minio設定
  // var Minio = require('minio')
  const minioClient = new minio.Client({
    endPoint: "play.min.io",
    port: 9000,
    useSSL: false,
    accessKey: "minio",
    secretKey: "minio123"
  });

  //バケット作成
  const createBucket = (bucketName) => {
    minioClient.makeBucket(bucketName, location, function(err) {
      if(err) return console.log(err);

      console.log('Bucket created successfully in ', location);
      console.log('Created bucket name  => ', bucketName);
    });
  }

  const metaData = {
    'Content-Type': 'application/octet-stream',
    'X-Amz-Meta-Testing': 1234,
    'example': 5678
  }


  const processImage = (event) => {
    const File = event.target.files[0];
    setImage(File);
    const imageUrl = URL.createObjectURL(File);
    setProfileImage(imageUrl);
    console.log("qqqqqq",File);
    const uploadURL =  BASEURL + File.name;
    setImageUploadUrl(uploadURL);
    console.log("aaaa",uploadURL);
  }

  //フォームsubmit
  const submitUserRegisterForm = async(event) => {
    event.preventDefault();
    //バケット作成
    // createBucket(BUCKET);
    //画像をminioに保存、画像ファイル名の取得
    console.log("aaaa",imageUploadUrl);
    //minioへのアップロード
    minioClient.fPutObject(BUCKET, imageFile.name, imageUploadUrl, metaData, (err, etag) => {
        if (err)
          return console.log(err);
        console.log('File upload successfully!!');
        console.log('Uploaded File Name => ', file);
      });

    const request = new CreateUserRequest();
    request.setName(name);
    request.setScore(score);
    //DBには画像のURLを保存(minio)
    request.setPhotoUrl(imageUploadUrl);
    const client = new HelloServiceClient("http://localhost:8080");
    const response = await client.createUser(request, {});
  }
  return (
    <div className="user-form-container">
      <h4>user-register-form</h4>
      <form onSubmit={submitUserRegisterForm}>
      {/* onSubmit={event => this.ubmitUserRegisterForm(event)} */}
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