// Copyright 2021 The Elabox Authors
// This file is part of elabox-logs library.

// elabox-logs library is under open source LGPL license.
// If you simply compile or link an LGPL-licensed library with your own code,
// you can release your application under any license you want, even a proprietary license.
// But if you modify the library or copy parts of it into your code,
// you’ll have to release your application under similar terms as the LGPL.
// Please check license description @ https://www.gnu.org/licenses/lgpl-3.0.txt


import { LOAD_FILTERS_AC, LOAD_LATEST_AC, DELETE_LOG_FILE_AC, PKID } from "../constant";

export function retrieveLatest(eventH, filter = {}, callback = (logs) => {}) {
    eventH.sendRPC(PKID, LOAD_LATEST_AC, '', filter)
        .then( logs => callback(logs))
}

export function retrieveSummary(eventH, callback) {
    eventH.sendRPC(PKID, LOAD_FILTERS_AC)
}

export function deleteLogFile(eventH, callback) {
    eventH.sendRPC(PKID, DELETE_LOG_FILE_AC)
        .then( res => {
            if (callback !== null)
                callback()
        })
}

