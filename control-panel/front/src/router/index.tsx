import * as React from 'react'
import * as ReactRouter from 'react-router-dom'

import { Component as Home } from './Home'

export const Route = (): React.ReactElement => (
    <ReactRouter.BrowserRouter>
        <ReactRouter.Routes>
            <ReactRouter.Route path="/">
                <ReactRouter.Route index element={<Home />} />
            </ReactRouter.Route>
        </ReactRouter.Routes>
    </ReactRouter.BrowserRouter>
)

export default Route
