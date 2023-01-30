export default [
  {
    url: `/mock/dashboard/home`,
    method: 'get',
    response: ({ params }) => {
      return {
        code: 200,
        data: {
          title: '我是首页'
        }
      }
    }
  }
]
