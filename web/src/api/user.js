import request from '@/utils/request';

export const login = data => {
    return request({
        url: 'grmapix/user/login',
        method: 'post',
        data
    });
};