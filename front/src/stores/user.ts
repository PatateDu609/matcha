import { defineStore } from 'pinia';
import { api } from 'boot/axios';

export type UserPayload = {
  'first-name': string;
  'last-name': string;
  username: string;
  'full-name': string;
  email: string;
  bio: string | undefined | null;
  score: number;
};

export type CurrentUserPayload = {
  id: string;
  'first-name': string;
  'last-name': string;
  username: string;
  'full-name': string;
  email: string;
  bio: string | undefined | null;
  score: number;
};

export interface UserState {
  uuid: string;
  firstName: string;
  lastName: string;
  username: string;
  fullName: string;
  email: string;
  bio: string | undefined | null;
  fameRating: number;
  loading: boolean;
}

function fillStore(store: UserState, data: CurrentUserPayload | UserPayload) {
  if ('id' in data) {
    store.uuid = data.id;
  }

  store.username = data.username;
  store.firstName = data['first-name'];
  store.lastName = data['last-name'];
  store.fullName = data['full-name'];
  store.email = data.email;
  store.bio = data.bio;
  store.fameRating = data.score;

  store.loading = false;
}

export const useUserStore = defineStore('user', {
  state: (): UserState => ({
    uuid: '',
    username: '',
    firstName: '',
    lastName: '',
    fullName: '',
    email: '',
    bio: '',
    fameRating: 0,

    loading: true,
  }),

  getters: {
    isLogged(): boolean {
      return !this.loading && this.uuid != '';
    },
  },

  actions: {
    async fetchUser(userID: string) {
      this.uuid = userID;

      await api
        .get<UserPayload>(`/user/${userID}`, {
          responseType: 'json',
        })
        .then((response) => {
          const data = response.data;

          fillStore(this, data);
        })
        .catch((reason) => {
          console.error(reason);
          this.loading = false;
        });
    },

    async fetchCurrentUser(): Promise<void> {
      return new Promise<void>((resolve, reject) => {
        api
          .get<CurrentUserPayload>('/user/', {
            responseType: 'json',
          })
          .then((response) => {
            fillStore(this, response.data);

            resolve();
          })
          .catch((reason) => {
            this.loading = false;
            reject(reason);
          });
      });
    },

    fillWithPayload(payload: UserPayload | CurrentUserPayload) {
      fillStore(this, payload);
    },
  },
});
