import { h, ref } from 'vue';
import { cloneDeep } from 'lodash-es';
import { FormSchema } from '@/components/Form';
import { defRangeShortcuts } from '@/utils/dateUtil';
import { renderImage, renderFile, renderOptionTag } from '@/utils';
import { useDictStore } from '@/store/modules/dict';

const dict = useDictStore();

export class State {
  public id = 0; // id
  public uuid = ''; // uuid
  public title = ''; // 标题
  public description = ''; // 摘要描述
  public coverUrl = ''; // 封面图
  public authorId = ''; // 作者
  public duration = 0; // 时长(s)
  public audioUrl = ''; // 音频
  public category = ''; // category
  public platform = ''; // platform
  public originUrl = ''; // 来源地址
  public content = ''; // 内容
  public scripts = []; // 脚本
  public zhubos = '{}'; // 主播配置
  public tags = []; // tags
  public auditStatus = null; // 审核状态
  public onlineStatus = null; // 上架状态
  public auditRemark = ''; // 审核备注
  public playCount = 0; // 播放次数
  public likeCount = 0; // 点赞次数
  public dislikeCount = 0; // 不喜欢次数
  public showCount = 0; // 展示次数
  public createdAt = ''; // 创建时间
  public updatedAt = ''; // 更新时间
  public deletedAt = ''; // 删除时间

  constructor(state?: Partial<State>) {
    if (state) {
      Object.assign(this, state);
    }
  }
}

export function newState(state: State | Record<string, any> | null): State {
  if (state !== null) {
    if (state instanceof State) {
      return cloneDeep(state);
    }
    return new State(state);
  }
  return new State();
}

// 表单验证规则
export const rules = {
  title: {
    required: true,
    trigger: ['blur', 'input'],
    type: 'string',
    message: '请输入标题',
  },
  description: {
    required: true,
    trigger: ['blur', 'input'],
    type: 'string',
    message: '请输入摘要描述',
  },
  coverUrl: {
    required: true,
    trigger: ['blur', 'input'],
    type: 'string',
    message: '请输入封面图',
  },
  authorId: {
    required: true,
    trigger: ['blur', 'input'],
    type: 'string',
    message: '请输入作者',
  },
  duration: {
    required: true,
    trigger: ['blur', 'input'],
    type: 'number',
    message: '请输入时长(s)',
  },
  audioUrl: {
    required: true,
    trigger: ['blur', 'input'],
    type: 'string',
    message: '请输入音频',
  },
  category: {
    required: true,
    trigger: ['blur', 'input'],
    type: 'string',
    message: '请输入category',
  },
  tags: {
    required: true,
    trigger: ['blur', 'input'],
    type: 'string',
    message: '请输入tags',
  },
};

// 表格搜索表单
export const schemas = ref<FormSchema[]>([
  {
    field: 'id',
    component: 'NInputNumber',
    label: 'id',
    componentProps: {
      placeholder: '请输入id',
      onUpdateValue: (e: any) => {
        console.log(e);
      },
    },
  },
  {
    field: 'uuid',
    component: 'NInput',
    label: 'uuid',
    componentProps: {
      placeholder: '请输入uuid',
      onUpdateValue: (e: any) => {
        console.log(e);
      },
    },
  },
  {
    field: 'auditStatus',
    component: 'NSelect',
    label: '审核状态',
    defaultValue: null,
    componentProps: {
      placeholder: '请选择审核状态',
      options: dict.getOption('sys_normal_disable'),
      onUpdateValue: (e: any) => {
        console.log(e);
      },
    },
  },
  {
    field: 'onlineStatus',
    component: 'NSelect',
    label: '上架状态',
    defaultValue: null,
    componentProps: {
      placeholder: '请选择上架状态',
      options: dict.getOption('sys_normal_disable'),
      onUpdateValue: (e: any) => {
        console.log(e);
      },
    },
  },
  {
    field: 'createdAt',
    component: 'NDatePicker',
    label: '创建时间',
    componentProps: {
      type: 'datetimerange',
      clearable: true,
      shortcuts: defRangeShortcuts(),
      onUpdateValue: (e: any) => {
        console.log(e);
      },
    },
  },
]);

// 表格列
export const columns = [
  {
    title: 'id',
    key: 'id',
    align: 'left',
    width: -1,
  },
  {
    title: '标题',
    key: 'title',
    align: 'left',
    width: -1,
  },
  {
    title: '摘要描述',
    key: 'description',
    align: 'left',
    width: -1,
  },
  {
    title: '封面图',
    key: 'coverUrl',
    align: 'left',
    width: -1,
    render(row: State) {
      return renderImage(row.coverUrl);
    },
  },
  {
    title: '作者',
    key: 'authorId',
    align: 'left',
    width: -1,
  },
  {
    title: '时长(s)',
    key: 'duration',
    align: 'left',
    width: -1,
  },
  {
    title: '音频',
    key: 'audioUrl',
    align: 'left',
    width: -1,
    render(row: State) {
      return renderFile(row.audioUrl);
    },
  },
  {
    title: 'category',
    key: 'category',
    align: 'left',
    width: -1,
  },
  {
    title: 'platform',
    key: 'platform',
    align: 'left',
    width: -1,
  },
  {
    title: '来源地址',
    key: 'originUrl',
    align: 'left',
    width: -1,
  },
  {
    title: '内容',
    key: 'content',
    align: 'left',
    width: -1,
  },
  {
    title: '脚本',
    key: 'scripts',
    align: 'left',
    width: -1,
  },
  {
    title: 'tags',
    key: 'tags',
    align: 'left',
    width: -1,
  },
  {
    title: '审核状态',
    key: 'auditStatus',
    align: 'left',
    width: -1,
    render(row: State) {
      return renderOptionTag('sys_normal_disable', row.auditStatus);
    },
  },
  {
    title: '上架状态',
    key: 'onlineStatus',
    align: 'left',
    width: -1,
    render(row: State) {
      return renderOptionTag('sys_normal_disable', row.onlineStatus);
    },
  },
  {
    title: '审核备注',
    key: 'auditRemark',
    align: 'left',
    width: -1,
  },
  {
    title: '播放次数',
    key: 'playCount',
    align: 'left',
    width: -1,
  },
  {
    title: '点赞次数',
    key: 'likeCount',
    align: 'left',
    width: -1,
  },
  {
    title: '不喜欢次数',
    key: 'dislikeCount',
    align: 'left',
    width: -1,
  },
  {
    title: '展示次数',
    key: 'showCount',
    align: 'left',
    width: -1,
  },
  {
    title: '创建时间',
    key: 'createdAt',
    align: 'left',
    width: -1,
  },
  {
    title: '更新时间',
    key: 'updatedAt',
    align: 'left',
    width: -1,
  },
];

// 加载字典数据选项
export function loadOptions() {
  dict.loadOptions(['sys_normal_disable']);
}