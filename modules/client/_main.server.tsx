import 'fast-text-encoding';
import React from 'react';
import { renderToString } from 'react-dom/server';
import { App, TAppProps } from '@modules/client/_main';
import RouterApp from '@modules/lib/react_router.figo';
import {
  createRoutesFromElements,
  matchRoutes,
} from 'react-router-dom';

type TRenderAppParms = TAppProps;
type TRenderAppReturn = { html: string };
export const renderApp = (
  arg: TRenderAppParms
): TRenderAppReturn => {
  const html = renderToString(<App url={arg.url} />);
  return { html };
};

export function getMatchRoutes(url: string) {
  const match = matchRoutes(
    createRoutesFromElements(<RouterApp />),
    url
  );
  return match;
}

// @ts-ignore
FIGOTSX.render = renderApp;
// @ts-ignore
FIGOTSX.getMatchRoutes = getMatchRoutes;
