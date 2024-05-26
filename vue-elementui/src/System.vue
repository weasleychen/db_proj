<template>
  <div>
    <el-container style="height: 100%; border: 1px solid #eee">
      <el-aside height=100% width="200px" style="background-color: white" class="el-aside">
        <el-menu :default-openeds="['1', '3']" class="left_menu">

          <el-submenu index="1">
            <template slot="title"><i class="el-icon-s-order"></i>桌台管理</template>
            <el-menu-item-group>
              <template slot="title">分组一</template>
              <el-menu-item index="1-1">
                <router-link to="/profile">个人信息</router-link>
                </el-menu-item>
              <el-menu-item index="1-2">
                <router-link to="/list">用户列表</router-link>
                </el-menu-item>
            </el-menu-item-group>
            <el-menu-item-group title="分组2">
              <el-menu-item index="1-3">选项3</el-menu-item>
            </el-menu-item-group>
            <el-submenu index="1-4">
              <template slot="title">选项4</template>
              <el-menu-item index="1-4-1">选项4-1</el-menu-item>
            </el-submenu>
          </el-submenu>
          <el-submenu index="2">
            <template slot="title"><i class="el-icon-fork-spoon"></i>就餐管理</template>
            <el-menu-item-group>
              <template slot="title">分组一</template>
              <el-menu-item index="2-1">选项1</el-menu-item>
              <el-menu-item index="2-2">选项2</el-menu-item>
            </el-menu-item-group>
            <el-menu-item-group title="分组2">
              <el-menu-item index="2-3">选项3</el-menu-item>
            </el-menu-item-group>
            <el-submenu index="2-4">
              <template slot="title">选项4</template>
              <el-menu-item index="2-4-1">选项4-1</el-menu-item>
            </el-submenu>
          </el-submenu>

          <el-submenu index="3">
            <template slot="title"><i class="el-icon-refrigerator"></i>菜单管理</template>
            <el-menu-item-group>
              <template slot="title">分组一</template>
              <el-menu-item index="3-1">选项1</el-menu-item>
              <el-menu-item index="3-2">选项2</el-menu-item>
            </el-menu-item-group>
            <el-menu-item-group title="分组2">
              <el-menu-item index="3-3">选项3</el-menu-item>
            </el-menu-item-group>
            <el-submenu index="3-4">
              <template slot="title">选项4</template>
              <el-menu-item index="3-4-1">选项4-1</el-menu-item>
            </el-submenu>
          </el-submenu>

          <el-submenu index="4">
            <template slot="title"><i class="el-icon-coordinate"></i>用户管理</template>
            <el-menu-item-group>
              <template slot="title">分组一</template>
              <el-menu-item index="3-1">选项1</el-menu-item>
              <el-menu-item index="3-2">选项2</el-menu-item>
            </el-menu-item-group>
            <el-menu-item-group title="分组2">
              <el-menu-item index="3-3">选项3</el-menu-item>
            </el-menu-item-group>
            <el-submenu index="3-4">
              <template slot="title">选项4</template>
              <el-menu-item index="3-4-1">选项4-1</el-menu-item>
            </el-submenu>
          </el-submenu>

          <el-submenu index="4">
            <template slot="title"><i class="el-icon-tickets"></i>营业额管理</template>
            <el-menu-item-group>
              <template slot="title">分组一</template>
              <el-menu-item index="3-1">选项1</el-menu-item>
              <el-menu-item index="3-2">选项2</el-menu-item>
            </el-menu-item-group>
            <el-menu-item-group title="分组2">
              <el-menu-item index="3-3">选项3</el-menu-item>
            </el-menu-item-group>
            <el-submenu index="3-4">
              <template slot="title">选项4</template>
              <el-menu-item index="3-4-1">选项4-1</el-menu-item>
            </el-submenu>
          </el-submenu>





        </el-menu>
      </el-aside>





      <el-container>
        <el-header style="text-align: right; font-size: 12px">
          <el-dropdown>
            <i class="el-icon-setting" style="margin-right: 15px"></i>
            <el-dropdown-menu slot="dropdown">
              <el-dropdown-item>查看</el-dropdown-item>
              <el-dropdown-item>新增</el-dropdown-item>
              <el-dropdown-item>删除</el-dropdown-item>
            </el-dropdown-menu>
          </el-dropdown>
          <span>邓宇飏</span>
        </el-header>

        <el-main>
<!--          <el-carousel indicator-position="outside">-->
<!--            <el-carousel-item>-->
<!--&lt;!&ndash;              <img src="../assets/backgg.jpg" alt="">&ndash;&gt;-->
<!--            </el-carousel-item>-->
<!--          </el-carousel>-->
          <router-view></router-view>

          <el-carousel :interval="4000" type="card" height="300px">
            <el-carousel-item v-for="item in imagebox" :key="item.id">
              <img :src="item.idView" class="image">
            </el-carousel-item>
          </el-carousel>

          <el-table :data="tableData">
            <el-table-column prop="date" label="姓名" width="140">
            </el-table-column>
            <el-table-column prop="name" label="感受" width="120">
            </el-table-column>
            <el-table-column prop="address" label="地址">
            </el-table-column>
          </el-table>

        </el-main>

      </el-container>
    </el-container>
  </div>

</template>


<style>

.el-header {
  background-color: #B3C0D1;
  color: #333;
  line-height: 60px;


}
.el-aside {
  color: red;
  width: 1000px;
  top:100px;

}
</style>

<script>
export default {
  data() {
    const item = {
      date: '2016-05-02',
      name: '王小虎',
      address: '上海市普陀区金沙江路 1518 弄'
    };
    return {
      imagebox:[
        {id:0,idView:require('./assets/backgg.jpg')},
        {id:1,idView:require('./assets/鹅肝.jpg')},
        {id:2,idView:require('./assets/牛肉.jpg')}
        //imagebox是assets下一个放图片的文件夹
      ],
      tableData: Array(1).fill(item)
    }

  }
};
</script>
