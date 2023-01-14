import { defineStore } from 'pinia';
import { api } from 'boot/axios';

type payload = {
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
  actions: {
    async fetchUser(userID: string) {
      this.uuid = userID;

      await api
        .get<payload>(`/user/${userID}`, {
          responseType: 'json',
        })
        .then((response) => {
          const data = response.data;

          console.log(data);

          this.username = data.username;
          this.firstName = data['first-name'];
          this.lastName = data['last-name'];
          this.fullName = data['full-name'];
          this.email = data.email;
          this.bio = data.bio;
          this.fameRating = data.score;

          this.loading = false;
        })
        .catch((reason) => {
          console.error(reason);
          this.loading = false;
        });
    },
  },
});
