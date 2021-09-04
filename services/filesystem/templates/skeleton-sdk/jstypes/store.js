import {EntityName} from "../apiModel";
import api from "../api";
import {findItemIndex} from "../common";

let findUrl = "/api/v1/entityName";
let readUrl = "/api/v1/entityName/"; // + id
let createUrl = "/api/v1/entityName";
let multiCreateUrl = "/api/v1/entityName/list";
let updateUrl = "/api/v1/entityName/"; // + id
let multiUpdateUrl = "/api/v1/entityName/list";
let deleteUrl = "/api/v1/entityName/"; // + id
let multiDeleteUrl = "/api/v1/entityName/list";
let findOrCreateUrl = "/api/v1/entityName";
let updateOrCreateUrl = "/api/v1/entityName";

const entityName = {
    actions: {
        createEntityName(context, {data, filter, header, noMutation}) {
            let url = createUrl;
            if (Array.isArray && Array.isArray(data)) {
                url = multiCreateUrl
            }
            return api.create(url, data, filter, header)
                .then(function (response) {
                    if (!noMutation) {
                        context.commit("setEntityName", response.Model);
                    }
                    return response;
                })
                .catch(function (err) {
                    console.error(err);
                    throw(err);
                });
        },
        deleteEntityName(context, {id, header, noMutation}) {
            let url;
            let dataOrNull = null;
            if (Array.isArray && Array.isArray(id)) {
                url = multiDeleteUrl;
                dataOrNull = id.map(item => typeof item === "number" ? {Id: item} : item);
            } else {
                url = deleteUrl + id;
            }
            return api.remove(url, header, dataOrNull)
                .then(function (response) {
                    if (!noMutation) {
                        context.commit("clearEntityName");
                    }
                    return response;
                })
                .catch(function (err) {
                    console.error(err);
                    throw(err);
                });
        },
        findEntityName(context, {filter, header, isAppend, noMutation}) {
            return api.find(findUrl, filter, header)
                .then(function (response) {
                    if (!noMutation) {
                        if (isAppend) {
                            context.commit("appendEntityName__List", response.List);
                        } else {
                            context.commit("setEntityName__List", response.List);
                        }
                    }
                    return response;
                })
                .catch(function (err) {
                    console.error(err);
                    throw(err);
                });
        },
        loadEntityName(context, {id, filter, header, noMutation}) {
            return api.find(readUrl + id, filter, header)
                .then(function (response) {
                    if (!noMutation) {
                        context.commit("setEntityName", response.Model);
                    }
                    return response;
                })
                .catch(function (err) {
                    console.error(err);
                    throw(err);
                });
        },
        updateEntityName(context, {id, data, filter, header, noMutation}) {
            let url = updateUrl + id;
            if (Array.isArray && Array.isArray(data)) {
                url = multiUpdateUrl
            }
            return api.update(url, data, filter, header)
                .then(function (response) {
                    if (!noMutation) {
                        context.commit("setEntityName", response.Model);
                    }
                    return response;
                })
                .catch(function (err) {
                    console.error(err);
                    throw(err);
                });
        },
        findOrCreateEntityName(context, {id, data, filter, header, noMutation}) {
            return api.update(findOrCreateUrl, data, filter, header)
                .then(function (response) {
                    if (!noMutation) {
                        context.commit("setEntityName", response.Model);
                    }
                    return response;
                })
                .catch(function (err) {
                    console.error(err);
                    throw(err);
                });
        },
        clearListEntityName(context) {
            context.commit("clearListEntityName");
        },
        clearEntityName(context) {
            context.commit("clearEntityName");
        },
    },
    getters: {
        getEntityName: (state) => {
            return state.EntityName;
        },
        getEntityNameById: state => id => {
            return state.EntityName__List.find(item => item.Id === id);
        },
        getListEntityName: (state) => {
            return state.EntityName__List;
        },
        getRoute__EntityName: state => action => {
            return state.EntityName__Routes[action];
        },
        getRoutes__EntityName: state => {
            return state.EntityName__Routes;
        },
    },
    mutations: {
        setEntityName(state, data) {
            state.EntityName = data;
        },
        setEntityName__List(state, data) {
            state.EntityName__List = data || [];
        },
        appendEntityName__List(state, data) {
            if (!state.EntityName__List) {
                state.EntityName__List = [];
            }
            if (data !== null) {
                state.EntityName__List = state.EntityName__List.concat(data);
            }
        },
        clearEntityName(state) {
            state.EntityName = new EntityName();
        },
        clearListEntityName(state) {
            state.EntityName__List = [];
        },
        updateEntityNameById(state, data) {
            let index = findItemIndex(state.EntityName__List, function (item) {
                return item.Id === data.Id;
            });

            if (index || index === 0) {
                state.EntityName__List.splice(index, 1, data);
            }
        },
        deleteEntityNameFromList(state, id) {
            let index = findItemIndex(state.EntityName__List, function (item) {
                return item.Id === id;
            });

            if (index || index === 0) {
                state.EntityName__List.splice(index, 1);
            }
        },
        addEntityNameItemToList(state, item) {
            if (state.EntityName__List === null) {
                state.EntityName__List = [];
            }
            state.EntityName__List.push(item);
        },
    },
    state: {
        EntityName: new EntityName(),
        EntityName__List: [],
        EntityName__Routes: {
            find: findUrl,
            read: readUrl,
            create: createUrl,
            multiCreate: multiCreateUrl,
            update: updateUrl,
            multiUpdate: multiUpdateUrl,
            delete: deleteUrl,
            multiDelete: multiDeleteUrl,
            findOrCreate: findOrCreateUrl,
            updateOrCreate: updateOrCreateUrl,
        },
    },
};

export default entityName;