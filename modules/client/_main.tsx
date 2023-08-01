import React from 'react';
import { StaticRouter } from 'react-router-dom/server';
import RouterApp from '@modules/lib/react_router.figo';
export type TAppProps = {
  url: string;
};

export const App: React.FC<TAppProps> = ({ url }) => {
  return (
    <StaticRouter location={url}>
      <RouterApp />
    </StaticRouter>
  );
};
