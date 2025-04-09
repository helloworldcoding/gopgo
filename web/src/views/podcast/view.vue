<template>
  <div>
    <n-drawer v-model:show="showModal" :width="dialogWidth">
      <n-drawer-content title="播客管理详情" closable>
        <n-spin :show="loading" description="请稍候...">
          <n-descriptions label-placement="left" class="py-2" :column="1">
            <n-descriptions-item>
              <template #label>
                标题
              </template>
              {{ formValue.title }}
            </n-descriptions-item>
            <n-descriptions-item>
              <template #label>
                摘要描述
              </template>
              {{ formValue.description }}
            </n-descriptions-item>
            <n-descriptions-item>
              <template #label>
                封面图
              </template>
              <n-image style="margin-left: 10px; height: 100px; width: 100px" :src="formValue.coverUrl"/>
            </n-descriptions-item>
            <n-descriptions-item>
              <template #label>
                作者
              </template>
              {{ formValue.authorId }}
            </n-descriptions-item>
            <n-descriptions-item>
              <template #label>
                时长(s)
              </template>
              {{ formValue.duration }}
            </n-descriptions-item>
            <n-descriptions-item>
              <template #label>
                音频
              </template>
              <div class="upload-card"  v-show="formValue.audioUrl !== ''" @click="download(formValue.audioUrl)">
                <div class="upload-card-item" style="height: 100px; width: 100px">
                  <div class="upload-card-item-info">
                    <div class="img-box">
                      <n-avatar :style="fileAvatarCSS">
                        {{ getFileExt(formValue.audioUrl) }}
                      </n-avatar>
                    </div>
                  </div>
                </div>
              </div>
            </n-descriptions-item>
            <n-descriptions-item>
              <template #label>
                category
              </template>
              {{ formValue.category }}
            </n-descriptions-item>
            <n-descriptions-item>
              <template #label>
                platform
              </template>
              {{ formValue.platform }}
            </n-descriptions-item>
            <n-descriptions-item>
              <template #label>
                来源地址
              </template>
              {{ formValue.originUrl }}
            </n-descriptions-item>
            <n-descriptions-item>
              <template #label>
                内容
              </template>
              <span v-html="formValue.content"></span>
            </n-descriptions-item>
            <n-descriptions-item>
              <template #label>
                脚本
              </template>
              {{ formValue.scripts }}
            </n-descriptions-item>
            <n-descriptions-item>
              <template #label>
                主播配置
              </template>
              {{ formValue.zhubos }}
            </n-descriptions-item>
            <n-descriptions-item>
              <template #label>
                tags
              </template>
              {{ formValue.tags }}
            </n-descriptions-item>
            <n-descriptions-item label="审核状态">
              <n-tag :type="dict.getType('sys_normal_disable', formValue.auditStatus)" size="small" class="min-left-space">
                {{ dict.getLabel('sys_normal_disable', formValue.auditStatus) }}
              </n-tag>
            </n-descriptions-item>
            <n-descriptions-item label="上架状态">
              <n-tag :type="dict.getType('sys_normal_disable', formValue.onlineStatus)" size="small" class="min-left-space">
                {{ dict.getLabel('sys_normal_disable', formValue.onlineStatus) }}
              </n-tag>
            </n-descriptions-item>
            <n-descriptions-item>
              <template #label>
                审核备注
              </template>
              {{ formValue.auditRemark }}
            </n-descriptions-item>
            <n-descriptions-item>
              <template #label>
                播放次数
              </template>
              {{ formValue.playCount }}
            </n-descriptions-item>
            <n-descriptions-item>
              <template #label>
                点赞次数
              </template>
              {{ formValue.likeCount }}
            </n-descriptions-item>
            <n-descriptions-item>
              <template #label>
                不喜欢次数
              </template>
              {{ formValue.dislikeCount }}
            </n-descriptions-item>
            <n-descriptions-item>
              <template #label>
                展示次数
              </template>
              {{ formValue.showCount }}
            </n-descriptions-item>
          </n-descriptions>
        </n-spin>
      </n-drawer-content>
    </n-drawer>
  </div>
</template>

<script lang="ts" setup>
  import { computed, ref } from 'vue';
  import { useMessage } from 'naive-ui';
  import { View } from '@/api/podcast';
  import { State, newState } from './model';
  import { adaModalWidth } from '@/utils/hotgo';
  import { getFileExt } from '@/utils/urlUtils';
  import { useDictStore } from '@/store/modules/dict';

  const message = useMessage();
  const dict = useDictStore();
  const loading = ref(false);
  const showModal = ref(false);
  const formValue = ref(newState(null));
  const dialogWidth = computed(() => {
    return adaModalWidth(580);
  });
  const fileAvatarCSS = computed(() => {
    return {
      '--n-merged-size': `var(--n-avatar-size-override, 80px)`,
      '--n-font-size': `18px`,
    };
  });

  // 下载
  function download(url: string) {
    window.open(url);
  }

  // 打开模态框
  function openModal(state: State) {
    showModal.value = true;
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

<style lang="less" scoped></style>