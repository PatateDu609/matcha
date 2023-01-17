import { defineStore } from 'pinia';

export interface RouterState {
  redirectURL: string | undefined;
}

export const useRouterStore = defineStore('router', {
  state: (): RouterState => ({
    redirectURL: undefined,
  }),
});
