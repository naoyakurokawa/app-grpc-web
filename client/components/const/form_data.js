export const FORMDATA  = [
  'DEFAULT',
];

export const STRAGEURL  = 'http://127.0.0.1:9000/';
export const BUCKET  = 'grpc-web-test';
export const BASEURL = STRAGEURL + BUCKET;

//テスト的にAPIのように、Languagesを取得する処理
export const getFormData = () =>{
  return new Promise((resolve) => {
    setTimeout( () => {
      resolve(FORMDATA);
    },1000)
  })
}