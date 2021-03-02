export const FORMDATA  = [
  'DEFAULT',
];

//テスト的にAPIのように、Languagesを取得する処理
export const getFormData = () =>{
  return new Promise((resolve) => {
    setTimeout( () => {
      resolve(FORMDATA);
    },1000)
  })
}