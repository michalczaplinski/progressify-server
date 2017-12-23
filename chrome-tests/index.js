const puppeteer = require('puppeteer');
const uniq = require('lodash').uniq;

const domains = [
  "facebook.com",
  "twitter.com",
  "google.com",
  "youtube.com",
  "linkedin.com",
  "wordpress.org",
  "instagram.com",
  "pinterest.com",
  "wikipedia.org",
  "wordpress.com",
  "blogspot.com",
  "adobe.com",
  "apple.com",
  "tumblr.com",
  "youtu.be",
  "amazon.com",
  "goo.gl",
  "vimeo.com",
  "microsoft.com",
  "flickr.com",
  "yahoo.com",
  "bit.ly",
  "buydomains.com",
  "qq.com",
  "godaddy.com",
  "vk.com",
  "reddit.com",
  "w3.org",
  "nytimes.com",
  "t.co",
];

(async () => {
  const browser = await puppeteer.launch();
  const page = await browser.newPage();
  await page.goto('http://gazeta.pl');
  
  const images = await page.evaluate(sel => document.querySelector('div'))
  console.log(images);

  images.forEach(async (image) => {

    console.log(image);
    // const result = await page.evaluate(async () => {
    //   let { status } = await fetch(`localhost:8081/${src}`);
    //   return status;
    // });

  });

  await browser.close();
})();