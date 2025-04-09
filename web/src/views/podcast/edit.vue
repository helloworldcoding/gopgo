<template>
  <div>
    <n-modal
      v-model:show="showModal"
      :mask-closable="false"
      :show-icon="false"
      preset="dialog"
      transform-origin="center"
      :title="formValue.id > 0 ? '编辑播客管理 #' + formValue.id : '添加播客管理'"
      :style="{
        width: dialogWidth,
      }"
    >
      <n-scrollbar style="max-height: 87vh" class="pr-5">
        <n-spin :show="loading" description="请稍候...">
          <n-form
            ref="formRef"
            :model="formValue"
            :rules="rules"
            :label-placement="settingStore.isMobile ? 'top' : 'left'"
            :label-width="100"
            class="py-4"
          >
            <n-grid cols="1 s:1 m:1 l:1 xl:1 2xl:1" responsive="screen">
              <n-gi span="1">
                <n-form-item label="标题" path="title">
                  <n-input placeholder="请输入标题" v-model:value="formValue.title" />
                </n-form-item>
              </n-gi>
              <n-gi span="1">
                <n-form-item label="摘要描述" path="description">
                  <n-input placeholder="请输入摘要描述" v-model:value="formValue.description" />
                </n-form-item>
              </n-gi>
              <n-gi span="1">
                <n-form-item label="封面图" path="coverUrl">
                  <UploadImage :maxNumber="1" v-model:value="formValue.coverUrl" />
                </n-form-item>
              </n-gi>
              <n-gi span="1">
                <n-form-item label="作者" path="authorId">
                  <n-input placeholder="请输入作者" v-model:value="formValue.authorId" />
                </n-form-item>
              </n-gi>
              <n-gi span="1">
                <n-form-item label="时长(s)" path="duration">
                  <n-input-number placeholder="请输入时长(s)" v-model:value="formValue.duration" />
                </n-form-item>
              </n-gi>
              <n-gi span="1">
                <n-form-item label="音频" path="audioUrl">
                  <UploadFile :maxNumber="1" v-model:value="formValue.audioUrl" />
                </n-form-item>
              </n-gi>
              <n-gi span="1">
                <n-form-item label="category" path="category">
                  <n-input placeholder="请输入category" v-model:value="formValue.category" />
                </n-form-item>
              </n-gi>
              <n-gi span="1">
                <n-form-item label="platform" path="platform">
                  <n-input placeholder="请输入platform" v-model:value="formValue.platform" />
                </n-form-item>
              </n-gi>
              <n-gi span="1">
                <n-form-item label="来源地址" path="originUrl">
                  <n-input placeholder="请输入来源地址" v-model:value="formValue.originUrl" />
                </n-form-item>
              </n-gi>
              <n-gi span="1">
                <n-form-item label="内容" path="content">
                  <Editor style="height: 450px" id="content" v-model:value="formValue.content" />
                </n-form-item>
              </n-gi>
              <n-gi span="1">
                <n-form-item label="脚本" path="scripts">
                  <n-input placeholder="请输入脚本" v-model:value="formValue.scripts" />
                </n-form-item>
              </n-gi>
              <n-gi span="1">
                <n-form-item label="主播配置" path="zhubos">
                  <n-input placeholder="请输入主播配置" v-model:value="formValue.zhubos" />
                </n-form-item>
              </n-gi>
              <n-gi span="1">
                <n-form-item label="tags" path="tags">
                  <n-input placeholder="请输入tags" v-model:value="formValue.tags" />
                </n-form-item>
              </n-gi>
              <n-gi span="1">
                <n-form-item label="审核状态" path="auditStatus">
                  <n-select v-model:value="formValue.auditStatus" :options="dict.getOptionUnRef('sys_normal_disable')" />
                </n-form-item>
              </n-gi>
              <n-gi span="1">
                <n-form-item label="上架状态" path="onlineStatus">
                  <n-select v-model:value="formValue.onlineStatus" :options="dict.getOptionUnRef('sys_normal_disable')" />
                </n-form-item>
              </n-gi>
              <n-gi span="1">
                <n-form-item label="审核备注" path="auditRemark">
                  <n-input placeholder="请输入审核备注" v-model:value="formValue.auditRemark" />
                </n-form-item>
              </n-gi>
              <n-gi span="1">
                <n-form-item label="播放次数" path="playCount">
                  <n-input-number placeholder="请输入播放次数" v-model:value="formValue.playCount" />
                </n-form-item>
              </n-gi>
              <n-gi span="1">
                <n-form-item label="点赞次数" path="likeCount">
                  <n-input-number placeholder="请输入点赞次数" v-model:value="formValue.likeCount" />
                </n-form-item>
              </n-gi>
              <n-gi span="1">
                <n-form-item label="不喜欢次数" path="dislikeCount">
                  <n-input-number placeholder="请输入不喜欢次数" v-model:value="formValue.dislikeCount" />
                </n-form-item>
              </n-gi>
              <n-gi span="1">
                <n-form-item label="展示次数" path="showCount">
                  <n-input-number placeholder="请输入展示次数" v-model:value="formValue.showCount" />
                </n-form-item>
              </n-gi>
            </n-grid>
          </n-form>
        </n-spin>
      </n-scrollbar>
      <template #action>
        <n-space>
          <n-button @click="closeForm">
            取消
          </n-button>
          <n-button type="info" :loading="formBtnLoading" @click="confirmForm">
            确定
          </n-button>
        </n-space>
      </template>
    </n-modal>
  </div>
</template>

<script lang="ts" setup>
  import { ref, computed } from 'vue';
  import { useDictStore } from '@/store/modules/dict';
  import { Edit, View } from '@/api/podcast';
  import { State, newState, rules } from './model';
  import UploadImage from '@/components/Upload/uploadImage.vue';
  import UploadFile from '@/components/Upload/uploadFile.vue';
  import Editor from '@/components/Editor/editor.vue';
  import { useProjectSettingStore } from '@/store/modules/projectSetting';
  import { useMessage } from 'naive-ui';
  import { adaModalWidth } from '@/utils/hotgo';

  const emit = defineEmits(['reloadTable']);
  const message = useMessage();
  const settingStore = useProjectSettingStore();
  const dict = useDictStore();
  const loading = ref(false);
  const showModal = ref(false);
  const formValue = ref<State>(newState(null));
  const formRef = ref<any>({});
  const formBtnLoading = ref(false);
  const dialogWidth = computed(() => {
    return adaModalWidth(840);
  });

  // 提交表单
  function confirmForm(e) {
    e.preventDefault();
    formRef.value.validate((errors) => {
      if (!errors) {
        formBtnLoading.value = true;
        Edit(formValue.value)
          .then((_res) => {
            message.success('操作成功');
            closeForm();
            emit('reloadTable');
          })
          .finally(() => {
            formBtnLoading.value = false;
          });
      } else {
        message.error('请填写完整信息');
      }
    });
  }

  // 关闭表单
  function closeForm() {
    showModal.value = false;
    loading.value = false;
  }

  // 打开模态框
  function openModal(state: State) {
    showModal.value = true;

    // 新增
    if (!state || state.id < 1) {
      formValue.value = newState(state);

      return;
    }

    // 编辑
    loading.value = true;
    View({ id: state.id })
      .then((res) => {
        formValue.value = res;
      })
      .finally(() => {
        loading.value = false;
      });
  }

  defineExpose({
    openModal,
  });
</script>

<style lang="less"></style>