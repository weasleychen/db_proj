<template>

<!--  <div id="login"><h1>首页</h1>-->
<!--    {{res}}-->
<!--  </div>-->

<!--  <div id="result">-->
<!--    <h1>首页</h1>-->
<!--    <ul>-->
<!--      <li v-for="item in items" :key="item.id">{{ item}}</li>-->
<!--    </ul>-->
<!--  </div>-->
  <div>
    <h2>Ping结果:</h2>
    <h1>后端接口数据</h1>

<!--    <ul>-->
<!--      <li v-for="data in backendData" :key="data.id">{{ data.name }}</li>-->
<!--    </ul>-->

  </div>
</template>


<script >
import axios from "axios";

export default {
  name: 'Main',
  mounted() {
    this.getFun()
  },
  methods: {
    async getFun() {
      const res = await axios.get('http://8.138.93.57:10000/ping?message=123')
      console.log('getFun', res)

    }
  },
};
</script>





<!--<script>-->
<!--import axios from 'axios';-->

<!--export default {-->
<!--  name: 'Main',-->
<!--  -->
<!--}-->

<!--axios.get('/api/ping?message=hey')-->
<!--  .then(response => {-->
<!--    // 处理响应数据-->
<!--    console.log(response.data)-->
<!--  })-->
<!--  .catch(error => {-->
<!--    // 处理错误-->
<!--    console.log(error)-->
<!--  })-->

<!--</script>-->



// const axios = require('axios');
// export default {
//   name: "Main"
// }
// const postFun = async (data) => {
//   const res = await axios.post('http://8.138.93.57:10000/add-table', data)
//   console.log('postFun', res)
// }

// export default {
//   data() {
//     return {
//       items: []
//     }
//   },
//   mounted() {
//     axios.get('http://8.138.93.57:10000/ping')
//       .then(response => {
//         this.items = response.data;
//       })
//       .catch(error => {
//         console.log(error);
//       });
//   }
// }
// export default {
//   data() {
//     return {
//       pingResult: ''
//     }
//   },
//   mounted() {
//     axios.get('http://8.138.93.57:10000/ping')
//       .then(response => {
//         this.pingResult = response.data
//       })
//       .catch(error => {
//         console.log(error)
//       })
//   }
// }


<style scoped>


</style>
