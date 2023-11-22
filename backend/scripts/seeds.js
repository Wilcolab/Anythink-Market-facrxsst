//TODO: seeds script should come here, so we'll be able to put some data in our local env
const mongoose = require('mongoose');

require('../models/User');
require('../models/Item');
require('../models/Comment');

const Item = mongoose.model('Item');
const Comment = mongoose.model('Comment');
const User = mongoose.model('User');

// Connect to MongoDB
if (process.env.MONGODB_URI) {
    mongoose.connect(process.env.MONGODB_URI);
} else {
    console.warn('Missing MONGODB_URI in the environment');
}

async function seedDatabase() {
    const users = Array.from(Array(100)).map((_item, i) => ({
        username: `fakeuser${i}`,
        email: `fakeuser${i}@anythinks.com`,
        bio: 'test bio',
        image: 'https://picsum.photos/200',
        role: 'user',
        favorites: [],
        following: [],
    }));

    for (let user of users) {
        const existingUser = await User.findOne({ username: user.username, email: user.email });

        if (!existingUser) {
            const u = new User(user);
            await u.save();
        }
    }

    const createdUsers = await User.find();

    const items = Array.from(Array(100)).map((_item, i) => ({
        slug: `fakeitem${i}`,
        title: `Fake Item${i}`,
        description: 'test description',
        image: 'https://picsum.photos/200',
        comments: [],
        tagList: ['test', 'tag'],
        seller: createdUsers[i]._id,
    }));

    for (let item of items) {
        const it = new Item(item);
        await it.save();
    }

    const createdItems = await Item.find();

    for (let i = 0; i < createdItems.length; i++) {
        const comments = Array.from(Array(100)).map((_item, j) => ({
            body: `This is the body ${j}`,
            seller: createdUsers[j]._id,
            item: createdItems[i]._id,
        }));

        for (let comment of comments) {
            const c = new Comment(comment);
            await c.save();
        }
    }
}

seedDatabase()
    .then(() => {
        console.log('Database seeded successfully.');
        process.exit(0);
    })
    .catch((err) => {
        console.error('Error seeding database:', err);
        process.exit(1);
    });
