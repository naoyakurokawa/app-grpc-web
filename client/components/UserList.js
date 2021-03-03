export const UserList = ({userList}) => {
  return(
    <div>
      {
        Object.keys(userList).map((key) => {
          return <div>名前：{userList[key].name} スコア：{userList[key].score}</div>
        })
      }
    </div>
  )
}