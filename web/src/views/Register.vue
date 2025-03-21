<template>
  <div>
    <div class="bg"></div>
    <div class="register-page">
      <div class="page-inner">
        <div class="contain" v-if="enableRegister">
          <div class="logo">
            <el-image src="images/logo.png" fit="cover"/>
          </div>

          <div class="header">{{ title }}</div>
          <div class="content">
            <el-form :model="formData" label-width="120px" ref="formRef">
              <div class="block">
                <el-input placeholder="请输入用户名(4-30位)"
                          size="large" maxlength="30"
                          v-model="formData.username"
                          autocomplete="off">
                  <template #prefix>
                    <el-icon>
                      <UserFilled/>
                    </el-icon>
                  </template>
                </el-input>
              </div>

              <div class="block">
                <el-input placeholder="请输入密码(8-16位)"
                          maxlength="16" size="large"
                          v-model="formData.password" show-password
                          autocomplete="off">
                  <template #prefix>
                    <el-icon>
                      <Lock/>
                    </el-icon>
                  </template>
                </el-input>
              </div>

              <div class="block">
                <el-input placeholder="重复密码(8-16位)"
                          size="large" maxlength="16" v-model="formData.repass" show-password
                          autocomplete="off">
                  <template #prefix>
                    <el-icon>
                      <Lock/>
                    </el-icon>
                  </template>
                </el-input>
              </div>

              <div class="block">
                <el-input placeholder="手机号码"
                          size="large" maxlength="11"
                          v-model="formData.mobile"
                          autocomplete="off">
                  <template #prefix>
                    <el-icon>
                      <Iphone/>
                    </el-icon>
                  </template>
                </el-input>
              </div>

              <div class="block" v-if="enableMsg">
                <el-row :gutter="10">
                  <el-col :span="12">
                    <el-input placeholder="手机验证码"
                              size="large" maxlength="30"
                              v-model.number="formData.code"
                              autocomplete="off">
                      <template #prefix>
                        <el-icon>
                          <Checked/>
                        </el-icon>
                      </template>
                    </el-input>
                  </el-col>
                  <el-col :span="12">
                    <send-msg size="large" :mobile="formData.mobile"/>
                  </el-col>
                </el-row>
              </div>

              <el-row class="btn-row">
                <el-button class="login-btn" size="large" type="primary" @click="register">注册</el-button>
              </el-row>

              <el-row class="text-line">
                已经有账号？
                <el-link type="primary" @click="router.push('login')">登录</el-link>
              </el-row>
            </el-form>
          </div>
        </div>

        <div class="tip-result" v-else>
          <el-result icon="error" title="注册功能已关闭">
            <template #sub-title>
              <p>抱歉，系统已关闭注册功能，请联系管理员添加账号！</p>
            </template>
          </el-result>
        </div>

        <footer class="footer">
          <footer-bar/>
        </footer>
      </div>
    </div>
  </div>
</template>

<script setup>

import {ref} from "vue";
import {Checked, Iphone, Lock, UserFilled} from "@element-plus/icons-vue";
import {httpGet, httpPost} from "@/utils/http";
import {ElMessage} from "element-plus";
import {useRouter} from "vue-router";
import FooterBar from "@/components/FooterBar.vue";
import SendMsg from "@/components/SendMsg.vue";

const router = useRouter();
const title = ref('ChatGPT-PLUS 用户注册');
const formData = ref({
  username: '',
  password: '',
  mobile: '',
  code: '',
  repass: '',
})
const formRef = ref(null)
const enableMsg = ref(false)
const enableRegister = ref(true)

httpGet('/api/sms/status').then(res => {
  if (res.data) {
    enableMsg.value = res.data['enabled_msg_service']
    enableRegister.value = res.data['enabled_register']
  }
})

const register = function () {
  if (formData.value.username.length < 4) {
    return ElMessage.error('用户名的长度为4-30个字符');
  }
  if (formData.value.password.length < 8) {
    return ElMessage.error('密码的长度为8-16个字符');
  }
  if (formData.value.repass !== formData.value.password) {
    return ElMessage.error('两次输入密码不一致');
  }

  if (formData.value.code === '') {
    formData.value.code = 0
  }
  httpPost('/api/user/register', formData.value).then(() => {
    ElMessage.success({"message": "注册成功，即将跳转到登录页...", onClose: () => router.push("login")})
  }).catch((e) => {
    ElMessage.error('注册失败，' + e.message)
  })
}

</script>

<style lang="stylus" scoped>
.bg {
  position fixed
  left 0
  right 0
  top 0
  bottom 0
  background-color #091519
  background-image url("~@/assets/img/reg-bg.jpg")
  background-size cover
  background-position center
  background-repeat no-repeat
  //filter: blur(10px); /* 调整模糊程度，可以根据需要修改值 */
}

.register-page {
  display flex
  justify-content center

  .page-inner {
    max-width 450px
    min-width 360px
    height 100vh
    display flex
    justify-content center
    align-items center

    .contain {
      padding 0 40px 20px 40px;
      width 100%
      color #ffffff
      border-radius 10px;
      z-index 10
      background-color rgba(255, 255, 255, 0.3)

      .logo {
        text-align center

        .el-image {
          width 120px;
        }
      }

      .header {
        width 100%
        margin-bottom 24px
        font-size 24px
        color $white_v1
        letter-space 2px
        text-align center
      }

      .content {
        width 100%
        height: auto
        border-radius 3px

        .block {
          margin-bottom 16px

          .el-input__inner {
            border 1px solid $gray-v6 !important

            .el-icon-user, .el-icon-lock {
              font-size 20px
            }
          }
        }

        .btn-row {
          padding-top 10px;

          .login-btn {
            width 100%
            font-size 16px
            letter-spacing 2px
          }
        }

        .text-line {
          justify-content center
          padding-top 10px;
          font-size 14px;
        }
      }
    }


    .tip-result {
      z-index 10
    }

    .footer {
      color #ffffff;

      .container {
        padding 20px;
      }
    }

  }

}
</style>

<style lang="stylus">
.register-page {
  .el-result {

    border-radius 10px;
    background-color rgba(14, 25, 30, 0.6)
    border 1px solid #666

    .el-result__title p {
      color #ffffff
    }

    .el-result__subtitle p {
      color #c1c1c1
    }
  }
}
</style>