import Link from 'next/link'
export const UserList = ({userList}) => {
  return(
    <div>
      {
        Object.keys(userList).map((key) => {
          return (
            <div>
              <a href={`/users/${userList[key].id}`}>{userList[key].id}</a>
              名前：{userList[key].name} スコア：{userList[key].score}
            </div>
          )
        })
      }
    </div>
  )
}