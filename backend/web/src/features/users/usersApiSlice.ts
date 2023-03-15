import { apiSlice } from "../../app/api/apiSlice";

export const usersApiSlice = apiSlice.injectEndpoints({
    endpoints: builder => ({
        getUsers: builder.query({ // this endpoint derive useGetUsersQuery hook
            query: () => '/users',
            keepUnusedDataFor: 5,
        }),
    })
})

export const {
    useGetUsersQuery
} = usersApiSlice
