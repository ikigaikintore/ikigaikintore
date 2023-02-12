# Control Panel

This app is for practice with React (maybe Redux too).

### Dependencies:



## Getting Started

First, run the development server:

```bash
yarn dev
```

Open [http://localhost:3000](http://localhost:3000) with your browser to see the result.

## CSS wrapping rules

When creating new components use this CSS wrapping rules:

```html
<body>
    <ComponentWrapper>
        <TitleWrapper>
            ...
        </TitleWrapper>
        <ContentWrapper>
            ...
        </ContentWrapper>
    </ComponentWrapper>
</body>
```

## Atomic design vs Design System Components

Atomic Design from [this article](https://medium.com/@janelle.wg/atomic-design-pattern-how-to-structure-your-react-application-2bb4d9ca5f97)  
Design System Component, recipes and snowflakes from [this article](https://medium.com/@janelle.wg/atomic-design-pattern-how-to-structure-your-react-application-2bb4d9ca5f97)

### Atomic Design

Atoms > Molecules > Organisms > Templates > Pages

- Atoms: buttons, input, form label...
- Molecules: Grouping Atoms
- Organisms: Combining molecules, navigation bar, cards...
- Templates: Groups of Organisms to form a page
- Pages: application

### Design System Components, recipes and snowflakes

- DSC: Shared components, content agnostic and context agnostic
- Recipes: Compositions of DSC living across a product
- Snowflakes: Components not reused in a product, they are unique

### Notes and blog contents

- https://zenn.dev/dove/articles/e940fa7e8b860d