import { Button, Dialog, Steps } from 'bkui-vue';
import { PropType, defineComponent, ref } from 'vue';
import './index.scss';
import AccountForm from './components/accountForm';
import AccountResource from './components/accountResource';
import ResultPage from './components/resultPage';

export default defineComponent({
  props: {
    isShow: {
      type: Boolean,
      required: true,
    },
    onSubmit: {
      type: Function as PropType<() => void>,
      required: true,
    },
    onCancel: {
      type: Function as PropType<() => void>,
      required: true,
    },
  },
  setup(props) {
    const step = ref(1);
    const enableNextStep = ref(false);
    const changeEnableNextStep = (val: boolean) => {
      enableNextStep.value = val;
    };
    return () => (
      <Dialog
        fullscreen
        isShow={props.isShow}
        class={'create-account-dialog-container'}>
        {{
          tools: () => (
            <div class={'create-account-dialog-tools'}>云账号接入</div>
          ),
          default: () => (
            <div class={'create-account-dialog-content'}>
              {step.value < 3 ? (
                <Steps
                  curStep={step.value}
                  class={'create-account-dialog-steps'}
                  steps={[
                    {
                      title: '录入账号',
                    },
                    {
                      title: '资源同步',
                    },
                  ]}
                />
              ) : (
                <ResultPage/>
              )}
              {step.value === 1 ? <AccountForm changeEnableNextStep={changeEnableNextStep}/> : null}
              {step.value === 2 ? <AccountResource /> : null}
            </div>
          ),
          footer: () => (
            <div class={'create-account-dialog-footer'}>
              {step.value > 1 ? (
                <Button
                  class={'mr8'}
                  onClick={() => (step.value -= 1)}>
                  上一步
                </Button>
              ) : null}
              {step.value < 2 ? (
                <Button
                  theme={'primary'}
                  class={'mr8'}
                  disabled={!enableNextStep.value}
                  onClick={() => (step.value += 1)}>
                  下一步
                </Button>
              ) : (
                <Button
                  theme={'primary'}
                  class={'mr8'}
                  onClick={() => (step.value += 1)}>
                  提交
                </Button>
              )}

              <Button onClick={props.onCancel}>取消</Button>
            </div>
          ),
        }}
      </Dialog>
    );
  },
});
