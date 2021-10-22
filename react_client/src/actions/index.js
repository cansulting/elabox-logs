import { CODE_SUCCESS, LOAD_FILTERS_AC, LOAD_LATEST_AC, PKID } from "../constant";

export function retrieveLatest(eventH, filter = {}, callback = (logs) => {}) {
    eventH.sendRPC(PKID, {id: LOAD_LATEST_AC, data: filter}, callback)
}

export function retrieveSummary(eventH, callback) {
    eventH.sendRPC(PKID, {id: LOAD_FILTERS_AC}, callback)
}