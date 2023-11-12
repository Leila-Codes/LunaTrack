import knex from 'knex';

const database = knex({
    client: 'pg',
    connection: {
        host: 'localhost',
        port: 5432,
        user: 'postgres',
        password: 'postgres',
        database: 'luna_track'
    }
});

async function init() {
    await database.schema.createTable('projects', (t) => {
        t.text('key').unique({indexName: 'idx_key'}).notNullable().primary();
        t.text('title').notNullable();
        t.text('description', 'longtext');
        t.bigint('owner_id').notNullable();
    })

    await database.schema.createTable('project_members', (t) => {
        t.text('key').references('projects.key');
        t.bigint('user_id'); // references lunadb.public.users
        t.primary(['key', 'user_id'])
    })

    await database.schema.createTable('issues', (t) => {
        t.text('key').unique().notNullable().primary()
        t.text('project').references('projects.key').notNullable();
        t.enum('issue_type', ['bug', 'task', 'enhancement', 'feature', 'user story']);
        t.text('summary');
        t.text('description', 'longtext');
        t.integer('status');
        t.timestamps();
    })

    await database.schema.createTable('comments', (t) => {
        t.bigIncrements('comment_id').primary();
        t.text('key').references('issues.key').notNullable();
        t.timestamps();
        t.integer('author');
        t.text('comment', 'longtext');
    })

    await database.schema.createTable('issue_links', (t) => {
        t.bigIncrements('link_id');
        t.text('parent_issues').notNullable().references('issues.key');
        t.text('child_issue').notNullable().references('issues.key');
        t.enum('link_type', ['blocks', 'refers', 'duplicates']).notNullable().defaultTo('refers');
    })

    await database.schema.createTable('issue_assignees', (t) => {
        t.text('issue_id').notNullable().references('issues.key');
        t.bigInteger('user_id').notNullable();
        t.enum('assignee_type', ['developer', 'assisting']).notNullable().defaultTo('developer')
    })

    process.exit(0);

}

init();